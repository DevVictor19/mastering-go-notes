package main

import (
	"fmt"
	"os"
	"strconv"
)

// os.Args slice store the arguments from a command line call
// flag package - parsing command line arguments
// os.Args is initialized by Go and is available to the program when referenced
// everything is consider as a string as os.Args is a slice of strings
func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}

	var min, max float64

	var initialized = 0
	for i := 1; i < len(args); i++ {
		n, err := strconv.ParseFloat(args[i], 64)
		if err != nil {
			continue
		}

		if initialized == 0 {
			min = n
			max = n
			initialized = 1
		}

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	fmt.Println("Max:", max)
	fmt.Println("Min:", min)
}
