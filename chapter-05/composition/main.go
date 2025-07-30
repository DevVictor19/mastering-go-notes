package main

import "fmt"

type A struct {
	F2 string
	F3 string
}

// composition
type B struct {
	F1 string
	A
}

func main() {
	test := B{F1: "test1", A: A{F2: "test1", F3: "test2"}}
	fmt.Println(test.F2) // access directly
}
