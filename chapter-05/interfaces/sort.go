package main

type Size struct {
	F1 int
	F2 string
	F3 int
}

type Person struct {
	F1 int
	F2 string
	F3 Size
}

// implements sort.Interface
type PersonSlice []Person

func (a PersonSlice) Len() int {
	return len(a)
}

func (a PersonSlice) Less(i, j int) bool {
	return a[i].F3.F1 < a[j].F3.F1
}

func (a PersonSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
