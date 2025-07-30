package main

import "fmt"

type Secret struct {
	Value string
}

type Entry struct {
	F1 int
	F2 string
	F3 Secret
}

func main() {
	A := Entry{100, "F2", Secret{"myPassword"}}
	testStruct(A)
	testStruct(A.F3)
	testStruct("A string")

	aInt := getInt()

	// safe way
	myInt, ok := aInt.(bool)
	if ok {
		fmt.Println("Type assertion successful", myInt)
	} else {
		fmt.Println("Type assertion failed")
	}

	// dangerous way
	myInt2 := aInt.(int)
	myInt2++

	// will panic
	_ = aInt.(bool)
}

func testStruct(x any) {
	// type switch
	switch T := x.(type) {
	case Secret:
		fmt.Println("Secret Type")
	case Entry:
		fmt.Println("Entry Type")
	default:
		fmt.Printf("Not supported type: %T\n", T)
	}
}

func getInt() any {
	return 1
}
