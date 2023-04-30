package main

import (
	"fmt"
	"bufio"
	"os"
	algo "gpt-chan/algorithm"
)

func main() {
	a := algo.NewAlgorithm()
	reader := bufio.NewScanner(os.Stdin)

	var text string	
	fmt.Print("\nEnter text: ")
	if reader.Scan() {
		text = reader.Text()
	}
	
	var pattern string
	fmt.Print("Enter pattern: ")
	if reader.Scan() {
		pattern = reader.Text()
	}

	res := a.KMP(text, pattern)
	switch res {
		case -1:
			fmt.Println("Pattern not found in text with KMP.")
		case -2:
			fmt.Println("Invalid input for KMP.")
		default:
			fmt.Printf("Pattern found at index %d with KMP.\n", res)
	}

	res = a.BM(text, pattern)
	switch res {
		case -1:
			fmt.Println("Pattern not found in text with BM.")
		case -2:
			fmt.Println("Invalid input for BM.")
		default:
			fmt.Printf("Pattern found at index %d with BM.\n", res)
	}
}