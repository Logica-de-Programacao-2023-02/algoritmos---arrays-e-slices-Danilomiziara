package main

import "fmt"

func main() {
	var tamanho int
	fmt.Print("Digite o tamanho do slice: ")
	fmt.Scan(&tamanho)
	numero := make([]int, tamanho)
	fmt.Print("Decida os valores dos ", tamanho, " elementos (somente inteiros): ")
	for i := 0; i < tamanho; i++ {
		fmt.Scan(&numero[i])
	}
	fmt.Println("Slice resultante:", numero)
	soma := 0
	for _, num := range numero {
		soma += num
	}
	fmt.Printf("Resultado da soma: %d\n", soma)
}
