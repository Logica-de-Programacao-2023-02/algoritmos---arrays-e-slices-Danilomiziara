package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}

	var i, j int
	fmt.Print("Informe o índice do primeiro elemento a ser trocado: ")
	fmt.Scan(&i)
	fmt.Print("Informe o índice do segundo elemento a ser trocado: ")
	fmt.Scan(&j)

	s[i], s[j] = s[j], s[i]

	fmt.Println("Slice resultante:", s)
}
