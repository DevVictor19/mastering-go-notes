package main

import (
	"fmt"
	"sort"
)

func main() {
	sInts := []int{1, 0, 2, -3, 4, 20}
	sFloats := []float64{1.0, 0.2, 0.22, -3, 4.1, -0.1}
	sStrings := []string{"aa", "a", "A", "Aa", "aab", "AAa"}

	sort.Ints(sInts)
	fmt.Println("sInts:", sInts)

	sort.Sort(sort.Reverse(sort.IntSlice(sInts)))
	fmt.Println("sInts reverse:", sInts)

	sort.Float64s(sFloats)
	fmt.Println("sFloats:", sFloats)

	sort.Sort(sort.Reverse(sort.Float64Slice(sFloats)))
	fmt.Println("sFloats reverse:", sFloats)

	sort.Strings(sStrings)
	fmt.Println("sStrings:", sStrings)

	sort.Sort(sort.Reverse(sort.StringSlice(sStrings)))
	fmt.Println("sStrings reverse:", sStrings)
}
