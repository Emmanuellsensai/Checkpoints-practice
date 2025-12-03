package piscine

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}

	for i := nb; i >= 2; i-- {
		isPrime := true

		for j := 2; j*j <= i; j++ {
			if i%j == 0 {

				isPrime = false
				break 
			}
		}
		if isPrime {
			return i
		}
	}

	return 0
}
