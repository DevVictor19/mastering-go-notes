package main

type numeric interface {
	int | int8 | int16 | int32 | int64 | float64
}

func add[T numeric](n1, n2 T) T {
	return n1 + n2
}
