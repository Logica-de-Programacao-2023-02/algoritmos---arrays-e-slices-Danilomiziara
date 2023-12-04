package main

import "fmt"

func main() {
	numero := []int{1, 2, 3, 4, 5}
	numero = append(numero[:2], numero[3:]...)
	fmt.Println(numero)
}
