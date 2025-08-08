package main

import "fmt"

func main() {
	n := 10
	i := funRet(n)
	j := funRet(-4)

	fmt.Println(i(4))
	fmt.Println(j(4))
}

func funRet(i int) func(int) int {
	if i < 0 {
		return func(k int) int {
			k = -k
			return k + k
		}
	}

	return func(k int) int {
		return k * k
	}
}
