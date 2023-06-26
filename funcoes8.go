package main

import (
	"errors"
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	p, err := pares(nums)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Números pares:", p)
	}
}

func pares(nums []int) ([]int, error) {
	if len(nums) == 0 {
		return nil, errors.New("Slice vazio")
	}

	p := []int{}
	for _, num := range nums {
		if num%2 == 0 {
			p = append(p, num)
		}
	}

	return p, nil
}
