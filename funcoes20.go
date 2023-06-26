package main

import (
	"errors"
	"fmt"
)

func main() {
	s := []string{"hello", "world", "go", "programming", "language"}
	long, err := Long(s)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings com mais de 5 caracteres:", long)
	}
}

func Long(s []string) ([]string, error) {
	if len(s) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	long := make([]string, 0)

	for _, s := range s {
		if len(s) > 5 {
			long = append(long, s)
		}
	}

	return long, nil
}
