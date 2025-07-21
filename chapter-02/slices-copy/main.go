package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3, 4, 5}
	a2 := []int{0, 0}
	a3 := []int{0, 0}
	//a5 := []int{10, 11, 12, 13, 14}

	copy(a2, a1) // destination, source

	fmt.Println("a2", a2)
	fmt.Println("a1", a1)

	copy(a1, a3) // destination, source

	fmt.Println("a1", a1)
	fmt.Println("a3", a3)
}
