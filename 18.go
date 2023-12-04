package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Digite a quantidade de números primos a serem gerados: ")
	fmt.Scan(&n)

	primes := []int{}
	i := 2
	for len(primes) < n {
		isPrime := true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
		i++
	}

	fmt.Println(primes)

}
