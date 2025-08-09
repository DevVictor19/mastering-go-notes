package main

import (
	"fmt"
	"os"
)

// os.O_APPEND - if file already exists, should append instead of truncating
// os.O_CREATE - if file does not already exists, it should be created
// os.O_WRONLY - the program should open the file for writing only

func main() {
	f, err := os.OpenFile("./test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	n, err := f.Write([]byte("Add data to file\n"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)

	n, err = f.Write([]byte("Add more data to file\n"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
}
