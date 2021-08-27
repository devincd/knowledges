// 修改 echo 程序，使其打印每个参数的索引和值，每个一行
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, val := range os.Args {
		fmt.Println("index:", index, "\t val:", val)
	}

	for val := range os.Args {
		fmt.Println(val)
	}
}
