package main

import (
	"fmt"
	// "bufio"
	// "os"
	algo "gpt-chan/algorithm"
)

func main() {
	// a := algo.New()
	// reader := bufio.NewScanner(os.Stdin)

	// var text string	
	// fmt.Print("\nEnter text: ")
	// if reader.Scan() {
	// 	text = reader.Text()
	// }
	
	// var pattern string
	// fmt.Print("Enter pattern: ")
	// if reader.Scan() {
	// 	pattern = reader.Text()
	// }

	// res := a.KMP(text, pattern)
	// switch res {
	// 	case -1:
	// 		fmt.Println("Pattern not found in text with KMP.")
	// 	case -2:
	// 		fmt.Println("Invalid input for KMP.")
	// 	default:
	// 		fmt.Printf("Pattern found at index %d with KMP.\n", res)
	// }

	// res = a.BM(text, pattern)
	// switch res {
	// 	case -1:
	// 		fmt.Println("Pattern not found in text with BM.")
	// 	case -2:
	// 		fmt.Println("Invalid input for BM.")
	// 	default:
	// 		fmt.Printf("Pattern found at index %d with BM.\n", res)
	// }
	
	// dateStr := "hari apa tanggal 30/05/2003, 20/05/2023"
	// dates := algo.ExtractDates(dateStr)

	// for _, date := range dates {
	// 	fmt.Printf("Date: %s\n", date)
	// 	day := algo.DateToDay(date)
	// 	fmt.Printf("Day: %s\n", day)
	// }

	addReq := "tambah pertanyaan siapa nama kamu dengan jawaban nama saya gpt-chan"
	fmt.Println("request: ", addReq)
	fmt.Println("ext. q?: ", algo.ExtractAddQuestions(addReq)[0])
	fmt.Println("ext. a?: ", algo.ExtractAnswers(addReq)[0])

	delReq := "hapus pertanyaan siapa nama kamu"
	fmt.Println("request: ", delReq)
	fmt.Println("ext. q?: ", algo.ExtractDeleteQuestions(delReq)[0])
}