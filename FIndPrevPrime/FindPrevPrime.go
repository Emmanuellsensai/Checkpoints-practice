package piscine

func FindPrevPrime(nb int) int {
	// 1. Edge Case: Numbers less than 2 cannot be prime.
	if nb < 2 {
		return 0
	}

	// 2. Loop Backwards from 'nb' down to 2.
	// We want the highest prime less than or equal to 'nb'.
	for i := nb; i >= 2; i-- {
		
		// Assume 'i' is prime until we find a divisor.
		isPrime := true

		// 3. Inner Loop: Check for divisors from 2 up to the square root of 'i'.
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				// Divisor found, 'i' is not prime.
				isPrime = false
				break
			}
		}

		// 4. If isPrime is still true, we found our number.
		if isPrime {
			return i
		}
	}

	return 0
}

