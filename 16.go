package main

import "fmt"

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var s []int
	for _, v := range a {
		if v%2 == 0 {
			s = append(s, v)
		}
	}

	fmt.Println("Novo Slice:", s)
}
