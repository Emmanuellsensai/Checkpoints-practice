package piscine

func RepeatAlpha(s string) string {
	result := " "
	for _, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			temp := int(ch) - 'A' + 1
			for i := 0; i < temp; i++ {
				result += string(rune(ch))
			}
		} else if ch >= 'a' && ch <= 'z' {
			temp := int(ch) - 'a' + 1
			for i := 0; i < temp; i++ {
				result += string(rune(ch))
			}
		} else {
			result += string(rune(ch))
		}
	}
	return result
}
