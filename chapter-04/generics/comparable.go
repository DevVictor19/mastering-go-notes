package main

func same[T comparable](v1, v2 T) bool {
	return v1 == v2
}
