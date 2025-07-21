package main

import (
	"fmt"
	"time"
)

func myPrint(start, finish int) {
	for i := start; i <= finish; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println(" ----- Start:", start, "Finish:", finish)
}

// go routines are initialized in a random order and start running in a random order
func main() {
	for i := range 4 {
		go myPrint(i, 5)
	}

	time.Sleep(time.Second)
}
