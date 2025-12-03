package piscine

import "github.com/01-edu/z01"

func PrintRevComb() {
	for a := 9; a >= 0; a-- {
		for b := a - 1; b >= 0; b-- {
			for c := b - 1; c >= 0; c-- {

				// Print the 3 digits
				z01.PrintRune(rune(a + '0'))
				z01.PrintRune(rune(b + '0'))
				z01.PrintRune(rune(c + '0'))

				// If this is NOT the last combination, print ", "
				if !(a == 2 && b == 1 && c == 0) {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}

	// Final newline
	z01.PrintRune('\n')
}

func main() {
	PrintRevComb()
}
