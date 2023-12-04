package main

import "fmt"

func main() {
	n := 3
	array1 := make([]int, n)
	array2 := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Print("Digite um número para o primeiro array: ")
		fmt.Scan(&array1[i])
	}
	for i := 0; i < n; i++ {
		fmt.Print("Digite um número para o segundo array: ")
		fmt.Scan(&array2[i])
	}

	sum := make([]int, n)
	for i := 0; i < n; i++ {
		sum[i] = array1[i] + array2[i]
	}

	fmt.Println(sum)

}
