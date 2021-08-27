package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// 低效版本
	countTime(printArgs1, os.Args, "printArgs1")
	// 高效版本
	countTime(printArgs2, os.Args, "printArgs2")
}

func countTime(f func([]string) string, params []string, funcName string) {
	now := time.Now().UnixNano()
	f(params)
	after := time.Now().UnixNano()
	fmt.Printf("func %s time cost %d ns\n", funcName, after-now)
}

// 低效版本
func printArgs1(args []string) string {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}

// 高效版本
func printArgs2(args []string) string {
	return strings.Join(args, " ")
}
