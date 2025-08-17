package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
	m  sync.Mutex
	v1 int = 1
)

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			duplicate()
			fmt.Printf("routine %d value %d -> ", i+1, read())
		}()
	}
	wg.Wait()
}

func duplicate() {
	m.Lock()
	defer m.Unlock()

	v1 = v1 * 2
}

func read() int {
	m.Lock()
	defer m.Unlock()

	return v1
}
