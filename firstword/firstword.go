package piscine

func FirstWord(s string) string {
	index := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			index = i
			break
		}
	}
	newLen := s[:index]
	return newLen + "\n"
}