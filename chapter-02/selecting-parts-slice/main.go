package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	length := len(s)
	fmt.Println("s:", s)
	fmt.Println("length:", length)
	fmt.Println("capacity:", cap(s))

	// first 5 elements
	fmt.Println(s[0:5])
	fmt.Println(s[:5])

	// last 2 elements
	fmt.Println(s[length-2 : length]) // último index não inclui
	fmt.Println(s[length-2:])

	s2 := s[0:5:10] // cap = 10 - 0 = 10
	fmt.Println(len(s2), cap(s2))

	s2 = s[2:5:10] // cap = 10 - 2 = 8
	fmt.Println(len(s2), cap(s2))

	s2 = s[:5:6] // cap = 6 - 0 = 6
	fmt.Println(len(s2), cap(s2))
}
