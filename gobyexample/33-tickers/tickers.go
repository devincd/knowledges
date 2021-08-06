/*
Timers are for when you want to do something once in the
future. Tickers are for when you want to do something
repeatedly at regular intervals. Here's an example of a
ticker that ticks periodically until we stop it.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// Tickers use a similar mechanism to timers: a channel that
	// is sent values. Here we'll use the select builtin on the
	// channel to await the values as they arrive every 500ms
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick as", t)
			}
		}
	}()

	// Tickers can be stopped like timers. Once a ticker is
	// stopped it won't receive any more values on its channel.
	// We'll stop ours after 1600ms.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

/*
OUTPUT
$ go run tickers.go
Tick as 2021-08-06 09:35:56.402493 +0800 CST m=+0.503973200
Tick as 2021-08-06 09:35:56.903735 +0800 CST m=+1.005225805
Tick as 2021-08-06 09:35:57.400428 +0800 CST m=+1.501928870
Ticker stopped
*/
