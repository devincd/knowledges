/*
Closing a channel indicates that no more values will be sent on it.
This can be useful to communicate completion to the channel's receivers.
*/
package main

import "fmt"

func main() {
	// In this example we'll use a jobs channel to communicate
	// work to be done from the main() goroutine to a worker
	// goroutine. When we have no more jobs for the worker
	// we'll close the jobs channel.
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Here's the worker goroutine. It repeatedly receives from
	// jobs with j, more := <- jobs. In this special 2-value
	// form of receive, the more value will be false if jobs has
	// been closed and all values in the channel have already
	// been received. We use this to notify on done when we've
	// worked all our jobs.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// This sends 3 jobs to the worker over the jobs channel,
	// the closes it.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}
	close(jobs)

	// We await the worker using the synchronization approach
	// we saw earlier
	fmt.Println("send all jobs")

	<-done
}

/*
OUTPUT
$ go run closing-channels.go
send job 1
send job 2
send job 3
send all jobs
received job 1
received job 2
received job 3
received all jobs
*/
