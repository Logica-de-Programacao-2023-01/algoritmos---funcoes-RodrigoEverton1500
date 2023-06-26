package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	input := "Hello, World!"
	resultado, err := re(input)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}

func re(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("A string está vazia")
	}

	v := "aeiouAEIOU"
	resultado := strings.Builder{}

	for _, c := range input {
		if strings.ContainsRune(v, c) {
			resultado.WriteRune('*')
		} else {
			resultado.WriteRune(c)
		}
	}

	return resultado.String(), nil
}
