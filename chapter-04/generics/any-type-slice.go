package main

import "fmt"

type slice[E any] interface {
	~[]E
}

func printLen[T slice[U], U any](s T) {
	fmt.Println(len(s))
}

// short cut (interface without enclosing keys)
func printLen2[S ~[]E, E any](x S) int {
	return len(x)
}

type numericSlice[E numeric] interface {
	~[]E
}

func doubleAll[T numericSlice[U], U numeric](s T) {
	for i, v := range s {
		s[i] = v * 2
	}
}
