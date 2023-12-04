package main

import "fmt"

func main() {
	numeros := [3]int{12, 2, 3}
	soma := 0
	for _, num := range numeros {
		soma += num
	}
	fmt.Printf("Resultado da soma: %d\n", soma)

}
