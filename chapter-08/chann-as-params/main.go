package main

import "fmt"

func main() {
	ch := make(chan bool)
	in(ch)
	out(ch)
}

// write only channel
func in(ch chan<- bool) {
	ch <- true
}

// readonly channel
func out(ch <-chan bool) {
	read := <-ch
	// ch <- true compiler error
	fmt.Println(read)
}
