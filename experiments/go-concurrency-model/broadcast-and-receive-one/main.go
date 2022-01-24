/*
Golang 并发模型-1
场景描述：广播一条消息给所有人，然后获取某一个人的正确反馈结果即可
*/
package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	go process2()

	result := process()
	fmt.Println(result)

	select {}
}

func process2() {
	for {
		fmt.Println(time.Now(), ":", runtime.NumGoroutine())
		time.Sleep(1 * time.Second)
	}
}

func process() string {
	targets := []string{"one", "two", "three", "four"}
	// channel 的大小应该为 targets 的长度，太小的话会出现协程泄漏的问题（协程阻塞地往 channel send 数据）
	resultChan := make(chan string, len(targets))
	for _, target := range targets {
		go func(targetStr string) {
			time.Sleep(5 * time.Second)
			// do something
			resultChan <- targetStr
		}(target)
	}

	select {
	case <-time.After(5 * time.Second):
		log.Println("timeout")
		return ""
	case result := <-resultChan:
		return result
	}
}
