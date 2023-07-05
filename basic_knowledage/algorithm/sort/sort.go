package main

import "fmt"

func main() {
	nums := []int{2, 5, 6, 3, 1, 4}

	fmt.Println(nums)
	bubbleSort(nums)
	fmt.Println("bubbleSort:", nums)
}

// 冒泡排序
// 原理：对给定的数组进行多次遍历，每次均比较相邻的两个数，如果前一个数比后一个数大，则交换这两个数，经过第一次遍历之后，最大的数就在
// 最右侧
func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		swap := false
		for j := 1; j < len(nums); j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}

// 插入排序
// 原理：从第二个数开始向右侧遍历，每次均把该位置的元素移动至左侧，放在一个正确的位置（比左侧大，比右侧小）
func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {

	}
}
