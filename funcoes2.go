package main

import (
	"fmt"
	"strings"
)

func main() {
	p := "Hello, World!"
	v := conta(p)
	fmt.Println("Quantidade de vogais:", v)
}

func conta(p string) int {
	v := []string{"a", "e", "i", "o", "u"}
	c := 0
	l := strings.ToLower(p)
	for _, char := range l {
		if cont(v, string(char)) {
			c++
		}
	}
	return c
}

func cont(slice []string, str string) bool {
	for _, val := range slice {
		if val == str {
			return true
		}
	}
	return false
}
