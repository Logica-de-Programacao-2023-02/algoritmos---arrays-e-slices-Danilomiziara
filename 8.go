package main

import (
	"fmt"
)

func main() {
	s := make([]string, 8)

	for i := range s {
		fmt.Printf("Digite o valor %d: ", i+1)
		fmt.Scan(&s[i])
	}

	var valor string
	fmt.Print("Digite o valor a ser removido: ")
	fmt.Scan(&valor)

	novoSlice := make([]string, 0, len(s))
	for _, v := range s {
		if v != valor {
			novoSlice = append(novoSlice, v)
		}
	}

	fmt.Println("Slice original:", s)
	fmt.Println("Slice sem o valor", valor, ":", novoSlice)
}
