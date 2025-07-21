package main

import "fmt"

// Reading from standard input
// fmt.Scanln can read user input while the program is already running
func main() {
	fmt.Printf("Please give me your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is", name)
}
