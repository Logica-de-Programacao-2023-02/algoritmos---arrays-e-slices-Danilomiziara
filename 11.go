package main

import (
	"fmt"
)

func main() {
	array := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	var linha, coluna int
	fmt.Print("Informe um índice de linha: ")
	fmt.Scanln(&linha)
	fmt.Print("Informe um índice de coluna: ")
	fmt.Scanln(&coluna)

	fmt.Printf("Valor armazenado na posição (%d, %d): %d\n", linha, coluna, array[linha][coluna])
}
