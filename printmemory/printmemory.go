package piscine

import "github.com/01-edu/z01"

func main() {
	piscine.PrintMemory([10]byte {'h', 'e', 'l', 'l', '1', '1', })
}

func PrintMemory(arr[10]byte) {
	length := len(arr)

	for i := 0; i < length; i++ {
		printHexByte(arr[i])
		if i == length-1 || i%4 == 3 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}
	for i := 0; i < length; i++ {
		b := arr[i]
		if b >= 32 && b <= 126 {
			z01.PrintRune(rune(b))
		}else {
			z01.PrintRune('.')
		}
		z01.PrintRune('\n')
		z01.PrintRune('\n')
	}
	func printHexByte(b byte) {
		high := b / 16
		low := b % 16
		z01.PrintRune(hexDigit(high))
		z01.PrintRune(hexDigit(low))
	}
	func hexDigit(n byte) rune {
		if n < 10 {
			return rune('0' + n)
		}
		return rune('a'+ (n-10))
	}
	
}