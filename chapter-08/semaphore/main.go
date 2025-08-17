package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

var (
	// number of active threads running goroutines
	workers = 4
	// no more than 4 workers can acquire the semaphore at the same time
	sem = semaphore.NewWeighted(int64(workers))
	wg  sync.WaitGroup
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need #jobs!")
		return
	}

	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	results := make([]int, nJobs)
	wg.Add(nJobs)

	ctx := context.TODO()
	for i := range results {
		fmt.Println("Acquire i:", i+1)
		if err := sem.Acquire(ctx, 1); err != nil {
			fmt.Println("Cannot acquire semaphore", err)
			break
		}

		go func(i int) {
			defer sem.Release(1)
			defer wg.Done()
			temp := worker(i)
			results[i] = temp
		}(i)
	}

	// if err := sem.Acquire(ctx, int64(workers)); err != nil {
	// 	fmt.Println(err)
	// }

	wg.Wait()

	for k, v := range results {
		fmt.Println(k, "->", v)
	}
}

func worker(n int) int {
	square := n * n
	time.Sleep(time.Second)
	return square
}
