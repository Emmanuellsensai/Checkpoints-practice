package main

import "fmt"

func FindPrevPrime(nb int) int {
	// Loop backwards from nb down to 2
	for i := nb; i >= 2; i-- {
		isPrime := true
		// Check divisors from 2 up to i-1 (regular check without sqrt optimization)
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		// If no divisors found, return immediately
		if isPrime {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}
