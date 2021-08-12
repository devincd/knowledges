/*
To wait for multiple goroutine finish, we can use a waitgroup
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

// This is the function we'll run in every goroutine. Note that
// a WaitGroup must be passed to function by pointer.
func worker(id int, wg *sync.WaitGroup) {

	// On return, notify the WaitGroup that we're done.
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	// Sleep to simulate an expensive task.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// This WaitGroup is used to wait for all the goroutines
	// launched here to finish.
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup
	// counter for each.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Block until the WaitGroup counter goes back to 0; all the
	// workers notified they're done.
	wg.Wait()
}

/*
OUTPUT
$ go run waitgroups.go
Worker 1 starting
Worker 2 starting
Worker 5 starting
Worker 4 starting
Worker 3 starting
Worker 3 done
Worker 4 done
Worker 2 done
Worker 5 done
Worker 1 done
*/
