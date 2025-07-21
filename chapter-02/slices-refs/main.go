package main

import "fmt"

func main() {
	ref1 := make([]int, 3) // length: 3, cap: 3
	ref2 := ref1

	ref1[0] = 1
	ref1[1] = 2
	ref1[2] = 3

	// aponta para o mesmo array interno
	fmt.Printf("ref1 %p\n", ref1) // ref1 0xa104040
	fmt.Printf("ref2 %p\n", ref2) // ref2 0xa104040

	ref2 = append(ref2, 4)         // capacidade do array vai ser modificada -> cap:6
	fmt.Println("ref2", cap(ref2)) // ref2 6

	// golang vai recriar um novo array interno pois a capacidade do antigo chegou no limite
	fmt.Printf("ref1 %p\n", ref1) // ref1 0xa104040
	fmt.Printf("ref2 %p\n", ref2) // ref2 0xa11a000
}
