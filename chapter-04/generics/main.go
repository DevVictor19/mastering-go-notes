package main

import "fmt"

func main() {
	fmt.Println(same("2", "4"))
	fmt.Println(same(2, 4))
	// fmt.Println(same("4", 4)) error

	fmt.Println(add(3, 4))
	fmt.Println(add(34.3, 3.4))
	fmt.Println(add(-3245.32, -43.32))
	// fmt.Println(add(-3245.32, 43)) error

	fmt.Println(addInt(2, 4))
	fmt.Println(addInt(-2, 4))

	var myInt1 customInt = 2
	var myInt2 customInt = 4

	// only allowed because i'm using the int supertype
	fmt.Println(addInt(myInt1, myInt2))

	s := []int{2, 3, 4}
	printLen(s)
	doubleAll(s)
	fmt.Println(s)
}
