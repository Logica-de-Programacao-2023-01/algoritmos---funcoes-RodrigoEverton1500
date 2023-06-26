package main

import "fmt"

func main() {
	nums := []int{5, 10, 2, 8, 3}
	val := 8
	index := Index(nums, val)
	fmt.Println("Index:", index)
}

func Index(nums []int, val int) int {
	for i, num := range nums {
		if num == val {
			return i
		}
	}
	return -1
}
