// This example demonstrates a priority queue built using the heap interface.
// source: https://golang.org/pkg/container/heap/
package util

import (
	"container/heap"
	db "gpt-chan/database/models"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    db.QA   // The value of the item; arbitrary.
	priority float64 // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func (item *Item) Value() db.QA {
	return item.value
}

func (item *Item) Priority() float64 {
	return item.priority
}

func (item *Item) Index() int {
	return item.index
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) PushVal(value db.QA, priority float64) {
	item := &Item{
		value:    value,
		priority: priority,
	}
	n := len(*pq)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) PriorityAt(i int) float64 {
	return (*pq)[i].priority
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value db.QA, priority float64) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func (pq *PriorityQueue) RemoveAt(i int) {
	heap.Remove(pq, i)
}

func (pq *PriorityQueue) PopVal() db.QA {
	return heap.Pop(pq).(*Item).value
}
