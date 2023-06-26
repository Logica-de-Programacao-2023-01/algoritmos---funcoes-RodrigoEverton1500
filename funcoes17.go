package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := []string{"banana", "maçã", "laranja", "abacaxi"}
	resultado, err := sorti(input)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}

func sorti(input []string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("O slice está vazio")
	}

	sort.Strings(input)
	resultado := strings.Join(input, ", ")

	return resultado, nil
}
