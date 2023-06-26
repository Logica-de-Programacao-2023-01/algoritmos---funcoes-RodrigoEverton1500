package main

import (
	"errors"
	"fmt"
)

func main() {
	num := 12345
	s, err := soma(num)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Soma dos dígitos de", num, ":", s)
	}
}

func soma(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("Número negativo")
	}

	s := 0
	for num > 0 {
		d := num % 10
		s += d
		num /= 10
	}

	return s, nil
}
