package main

import "fmt"

type User struct {
	ID   string
	Name string
	Role Role
}

type Role int

const (
	Admin Role = iota * 2 // you can use expressions
	Teacher
	_ // skip sequence
	Student
)

func main() {
	fmt.Println("Admin:", Admin)
	fmt.Println("Teacher:", Teacher)
	fmt.Println("Student:", Student)
}
