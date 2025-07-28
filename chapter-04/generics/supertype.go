package main

type customInt int

type allInts interface {
	~int // supertype (all ints included)
}

func addInt[T allInts](i1, i2 T) T {
	return i1 + i2
}
