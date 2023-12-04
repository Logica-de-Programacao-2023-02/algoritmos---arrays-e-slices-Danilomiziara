package main

import (
	"fmt"
	"math"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var min, max int = math.MaxInt32, math.MinInt32
	for _, valor := range s {
		if valor < min {
			min = valor
		}
		if valor > max {
			max = valor
		}
	}

	fmt.Println("Slice:", s)
	fmt.Println("Valor mínimo:", min)
	fmt.Println("Valor máximo:", max)
}
