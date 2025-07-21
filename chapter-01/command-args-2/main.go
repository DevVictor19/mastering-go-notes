package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Not enough arguments")
		return
	}

	var total, nInts, nFloats int

	invalid := make([]string, 0)

	for _, arg := range args[1:] {
		if _, err := strconv.Atoi(arg); err == nil {
			total++
			nInts++
			continue
		}

		if _, err := strconv.ParseFloat(arg, 64); err == nil {
			total++
			nFloats++
			continue
		}

		invalid = append(invalid, arg)
	}

	fmt.Println("#read:", total, "#ints", nInts, "#floats", nFloats)
	if len(invalid) > total {
		fmt.Println("Too much invalid input:", len(invalid))
		for _, s := range invalid {
			fmt.Println(s)
		}
	}
}
