package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need a delay!")
		return
	}

	delay, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Delay:", delay)
	f1(delay)
	f2(delay)
	f3(delay)
}

// The f1() function creates and executes a goroutine.
//
// If the c1 context calls the Done() function in less than 4 seconds,
// the goroutine will not have enough time to finish.
func f1(t int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c1.Done():
		fmt.Println("f1() Done:", c1.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f1() Interrupted:", r)
	}
}

// The f2() function creates and executes a goroutine.
//
// context.WithTimeout() requires two parameters a ctx and a timeout.
// When the timeout expires, the cancel function is called automatically.
func f2(t int) {
	c2 := context.Background()
	c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("f2() Done:", c2.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f2() Interrupted:", r)
	}
}

// The f3() function creates and executes a goroutine.
//
// context.WithDeadline() requires two parameters a ctx and a time in future (deadline).
// When the deadline passes the cancel function is called automatically.
func f3(t int) {
	c3 := context.Background()
	deadline := time.Now().Add(time.Duration(t) * time.Second)
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c3.Done():
		fmt.Println("f3() Done:", c3.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f3() Interrupted:", r)
	}
}
