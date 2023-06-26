package main

import (
	"fmt"
)

func main() {
	nums := []int{10, 20, 30, 40, 50}
	media := calculateAverage(nums)
	fmt.Println("Média:", media)
}

func calculateAverage(nums []int) float64 {
	soma := 0
	for _, num := range nums {
		soma += num
	}
	media := float64(soma) / float64(len(nums))
	return media
}
