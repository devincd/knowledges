package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var numb uint32 = 1

	fmt.Println(atomic.CompareAndSwapUint32(&numb, 1, 2))
	fmt.Println(numb)
}
