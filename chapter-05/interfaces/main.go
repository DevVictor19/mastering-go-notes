package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []Person{
		Person{1, "One", Size{1, "Person_1", 10}},
		Person{2, "Two", Size{2, "Person_2", 20}},
		Person{-1, "Two", Size{-1, "Person_3", -20}},
	}

	fmt.Println("Before", data)
	sort.Sort(PersonSlice(data))
	fmt.Println("After", data)

	sort.Sort(sort.Reverse(PersonSlice(data)))
	fmt.Println("Reverse", data)
}
