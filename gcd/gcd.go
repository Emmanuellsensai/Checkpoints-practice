package piscine

func Gcd(a, b unit) unit {
	if a == 0 || b == 0 {
		return 0
	}
	for b != 0 {
		div := a % b
		a = b
		b = div
	}
	return a
}
