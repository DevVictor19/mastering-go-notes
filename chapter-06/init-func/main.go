package main

import (
	"fmt"

	"github.com/DevVictor19/mastering-go-notes/chapter-06/init-func/p1"
)

func init() {
	fmt.Println("main package: init func executed")
}

func main() {
	p1.Fn1()
}
