package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	text := "Olá, mundo! Esta é uma string de exemplo."
	p, err := palavras(text)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Palavras:", p)
	}
}

func palavras(text string) ([]string, error) {
	if text == "" {
		return nil, errors.New("String vazia")
	}

	p := strings.Split(text, " ")
	return p, nil
}
