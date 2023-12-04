package main

import "fmt"

func main() {
	a := [10]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}

	var s []float64
	for _, valor := range a {
		if valor > 5 {
			s = append(s, valor)
		}
	}

	fmt.Println("Novo Slice:", s)
}
