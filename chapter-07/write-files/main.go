package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// example 2
	f1, err := os.Create("./f1.txt") // if already exists, truncates it
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f1.Close()
	fmt.Fprintf(f1, "Data to write\n")

	// example 1
	f2, err := os.Create("./f2.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f2.Close()
	n, err := f2.WriteString("Write the data")
	fmt.Printf("wrote %d bytes\n", n)

	// example 3
	f3, err := os.Create("./f3.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f3.Close()
	w := bufio.NewWriter(f3)
	n, err = w.WriteString("Write string")
	fmt.Printf("wrote %d bytes\n", n)
	w.Flush()
}
