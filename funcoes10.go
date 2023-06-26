package main

import (
	"errors"
	"fmt"
	"sort"
)

func main() {
	nums := []int{9, 3, 7, 1, 5}
	sorti, err := sortSlice(nums)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Números ordenados:", sorti)
	}
}

func sortSlice(nums []int) ([]int, error) {
	if len(nums) == 0 {
		return nil, errors.New("Slice vazio")
	}

	sorti := make([]int, len(nums))
	copy(sorti, nums)
	sort.Ints(sorti)
	return sorti, nil
}
