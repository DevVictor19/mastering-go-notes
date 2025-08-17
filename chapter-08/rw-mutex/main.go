package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	Password secret
	wg       sync.WaitGroup
)

type secret struct {
	RWM      sync.RWMutex
	password string
}

func main() {
	wg.Add(4)

	go show()
	go show()
	go show()
	go change("user")

	wg.Wait()
}

func change(pass string) {
	defer wg.Done()
	fmt.Println("Change() function")
	Password.RWM.Lock()
	fmt.Println("Change() Locked")
	time.Sleep(4 * time.Second)
	Password.password = pass
	Password.RWM.Unlock()
	fmt.Println("Change() Unlocked")
}

func show() {
	defer wg.Done()
	Password.RWM.RLock()
	fmt.Println("Show function locked")
	time.Sleep(2 * time.Second)
	fmt.Println("Pass value:", Password.password)
	Password.RWM.RUnlock()
}
