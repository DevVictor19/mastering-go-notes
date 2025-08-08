package p1

import (
	"fmt"

	"github.com/DevVictor19/mastering-go-notes/chapter-06/init-func/p2"
)

func init() {
	fmt.Println("p1 package: init func executed")
}

func Fn1() {
	fmt.Println("p1 package: execute Fn1")
	p2.Fn2()
}
