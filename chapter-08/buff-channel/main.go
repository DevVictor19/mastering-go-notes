package main

import "fmt"

func main() {
	requests := make(chan int, 5)

	for i := 0; i < 10; i++ {
		select {
		case requests <- i * i:
			fmt.Println("About to process", i)
		default:
			fmt.Println("No space for", i)
		}
	}

	fmt.Println()

	for {
		select {
		case num := <-requests:
			fmt.Println("*", num)
		default:
			fmt.Println("Nothing more to read")
			return
		}
	}
}
