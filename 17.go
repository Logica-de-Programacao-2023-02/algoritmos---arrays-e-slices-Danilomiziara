package main

import "fmt"

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var sum int
	for i, v := range a {
		if i%2 == 0 {
			sum += v
		}
	}

	fmt.Println("Soma dos elementos nas posições pares:", sum)
}
