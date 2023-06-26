package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	num := 17
	primo, err := check(num)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println(num, "é primo?", primo)
	}
}

func check(num int) (bool, error) {
	if num < 0 {
		return false, errors.New("Número negativo")
	}

	if num < 2 {
		return false, nil
	}

	sqrt := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqrt; i++ {
		if num%i == 0 {
			return false, nil
		}
	}

	return true, nil
}
