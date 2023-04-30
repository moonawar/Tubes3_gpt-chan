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
			fmt.Println("Pattern not found in text.")
		case -2:
			fmt.Println("Invalid input.")
		default:
			fmt.Printf("Pattern found at index %d.\n", res)
	}
}