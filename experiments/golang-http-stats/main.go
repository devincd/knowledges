/*

 */
package main

import (
	"fmt"
	"math"
)

func main() {
	v := math.Log(-1)
	m := map[float64]int{v: 1, v: 2, v: 3}
	fmt.Println(m[v], len(m))
}
