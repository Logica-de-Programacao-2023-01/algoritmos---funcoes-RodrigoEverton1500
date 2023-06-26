package main

import "fmt"

func main() {
	str := []string{"Hello", "World", "!"}
	conca := concat(str)
	fmt.Println("Concatenated string:", conca)
}

func concat(strings []string) string {
	resultado := ""
	for _, str := range strings {
		resultado += str
	}
	return resultado
}
