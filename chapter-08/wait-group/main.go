package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// negativeSyncWg()
	deadlockSyncWg()
}

// Call Add() less than Done()
//
// panic: sync: negative wait group counter
func negativeSyncWg() {
	var wg sync.WaitGroup

	count := 5
	wg.Add(1)

	for i := 0; i < count; i++ {

		go func(n int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println("Goroutine", i+1)
		}(i, &wg)
	}

	time.Sleep(200 * time.Millisecond) // wait to all goroutines call wg.Done()
	wg.Wait()
}

// call Add() more than Done()
//
// fatal error: all goroutines are asleep - deadlock
func deadlockSyncWg() {
	var wg sync.WaitGroup

	count := 5
	wg.Add(6)

	for i := 0; i < count; i++ {

		go func(n int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println("Goroutine", i+1)
		}(i, &wg)
	}

	time.Sleep(200 * time.Millisecond) // wait to all goroutines call wg.Done()
	wg.Wait()
}
