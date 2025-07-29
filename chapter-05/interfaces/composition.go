package main

// composition
type ABC interface {
	A
	B
	C
}

type A interface {
	MethodA() int
}

type B interface {
	MethodB() int
}

type C interface {
	MethodC() int
}
