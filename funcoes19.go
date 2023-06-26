package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	primos, err := Primos(20)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Números primos:", primos)
	}
}

func Primos(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("O número é negativo")
	}

	primos := make([]int, 0)

	for i := 2; i <= n; i++ {
		if Primo(i) {
			primos = append(primos, i)
		}
	}

	return primos, nil
}

func Primo(num int) bool {
	if num < 2 {
		return false
	}

	r := int(math.Sqrt(float64(num)))

	for i := 2; i <= r; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
