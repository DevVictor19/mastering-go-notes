package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type userTable struct {
	userSk int64
}

func (u *userTable) seq() int64 {
	return atomic.LoadInt64(&u.userSk)
}

func (u *userTable) increment() {
	atomic.AddInt64(&u.userSk, 1)
}

func main() {
	x := 100
	y := 4

	var wg sync.WaitGroup
	userT := userTable{}

	for i := 0; i < x; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < y; i++ {
				userT.increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println("current sequence:", userT.seq())
}
