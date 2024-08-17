package main

import (
	"fmt"
)

func generatePrimes(n int) []int {
	primes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if primes[i] {
			for j := i * i; j <= n; j += i {
				primes[j] = false
			}
		}
	}

	var result []int
	for i := 2; i <= n; i++ {
		if primes[i] {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	n := 100
	primes := generatePrimes(n)
	fmt.Printf("Nombres premiers jusqu'à %d: %v\n", n, primes)
}
