package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	str := []string{"Olá", "Mundo", "Golang"}
	resultado, err := concat(str)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}

func concat(str []string) (string, error) {
	if len(str) == 0 {
		return "", errors.New("Slice vazio")
	}
	return strings.Join(str, ","), nil
}
