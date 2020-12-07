package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	// 非range遍历不会出现这个问题
	//for i := 0; i < len(slice); i++ {
	//	m[i] = &slice[i]
	//}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
