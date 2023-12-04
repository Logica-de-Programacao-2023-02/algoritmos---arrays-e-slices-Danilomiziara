package main

import "fmt"

func main() {
	matriz := [3][2]int{}
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			fmt.Printf("Digite os valores de cada elemento da matriz: [%d] [%d]\n", i, j)
			fmt.Scan(&matriz[i][j])
		}
	}

	fmt.Printf("Matriz resultante:\n")
	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			fmt.Printf("%d\t", matriz[i][j])
		}
		fmt.Println()
	}
}
