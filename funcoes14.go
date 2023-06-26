package main

import (
	"errors"
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{3, 4, 5, 6}
	inter, err := interse(slice1, slice2)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Interseção:", inter)
	}
}

func interse(slice1, slice2 []int) ([]int, error) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return nil, errors.New("Um dos slices está vazio")
	}

	set := make(map[int]bool)
	inter := []int{}

	for _, num := range slice1 {
		set[num] = true
	}

	for _, num := range slice2 {
		if set[num] {
			inter = append(inter, num)
		}
	}

	return inter, nil
}
