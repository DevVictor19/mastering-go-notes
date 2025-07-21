package main

import "fmt"

func main() {
	s1 := []int{}
	s2 := []int{}

	s2 = append(s1, 1, 2, 3)

	fmt.Println(s1)
	fmt.Println(s2)
}
