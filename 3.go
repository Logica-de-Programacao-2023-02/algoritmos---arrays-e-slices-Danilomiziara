package main

import "fmt"

func main() {
	numero := [4]float64{50.2, 22.2, 30.6, 123.4}
	soma := 0.0
	for _, num := range numero {
		soma += num
	}
	fmt.Println("Resultado da soma: ", soma)
}
