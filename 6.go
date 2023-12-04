package main

import "fmt"

func main() {
	numeros := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var Valor int
	fmt.Print("Digite um valor inteiro:")
	fmt.Scan(&Valor)
	presente := false
	for _, num := range numeros {
		if num == Valor {
			presente = true
			break
		}
	}
	if presente == true {
		fmt.Println("Valor esta no array")
	} else {
		fmt.Println("Valor NAO esta no array")
	}
}
