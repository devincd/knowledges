/*
We often want to execute Go codes at some point in the
future, or repeatedly at some interval. Go's built-in timer
and ticker features make both of these tasks easy. We'll
look first at timers and then at tickers.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// Timers represent a single event in the future. You tell the
	// timer how long you want to wait, and it provides a channel
	// that will be notified at that time. This timer will wait 2
	// seconds.
	time1 := time.NewTimer(2 * time.Second)

	// The <-time1.C blocks on the timer's channel C until it
	// sends a value indicating that the timer fired.
	<-time1.C
	fmt.Println("Time 1 fired")

	// If you just wanted to wait, you could have used
	// time.Sleep. One reason a timer may be useful is that you
	// can cancel the timer before it fires. Here's an example of
	// that.
	time2 := time.NewTimer(time.Second)
	go func() {
		<-time2.C
		fmt.Println("Time 2 fired")
	}()

	stop2 := time2.Stop()
	if stop2 {
		fmt.Println("Time 2 stopped")
	}

	// Give the timer 2 enough time to fire, if it even was going to,
	// to show it is in fact stopped.
	time.Sleep(2 * time.Second)
}

/*
OUTPUT
$ go run timers.go
Time 1 fired
Time 2 stopped
*/
