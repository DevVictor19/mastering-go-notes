package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("You are using ", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("Using Go version", runtime.Version())
	fmt.Println("GOXMAXPROCS", runtime.GOMAXPROCS(0))

	go test(10)
}

func test(n int) int {
	fmt.Println(n)
	return n
}
