package main

import (
	"errors"
	"fmt"
)

func main() {
	num := 3
	slice := []int{1, 2, 3, 4}
	f, err := cont(num, slice)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Número encontrado:", f)
	}
}

func cont(num int, slice []int) (bool, error) {
	if len(slice) == 0 {
		return false, errors.New("O slice está vazio")
	}

	for _, val := range slice {
		if val == num {
			return true, nil
		}
	}

	return false, nil
}
