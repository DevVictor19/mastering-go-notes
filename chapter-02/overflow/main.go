package main

import (
	"fmt"
	"math"
)

func main() {
	i := 0

	for {
		fmt.Println("i:", i)
		fmt.Println("MaxInt:", math.MaxInt)

		if i == math.MaxInt {
			break
		}

		i = i + 1
	}

	fmt.Println("Max int:", i)
}
