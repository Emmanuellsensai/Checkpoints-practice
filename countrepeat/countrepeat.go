package piscine

func CountRepeat(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			count++
		}
	}
	return count
}
