package piscine

func HashCode(dec string) string {
	//result := ""
	size := len(dec)
	if size == 0 {
		return ""
	}
	var result string
	for _, ch := range dec {
		ascii := int(ch)
		hashed := (ascii + size) % 127
		if hashed > 127 || hashed < 0 {
			hashed += 33
		}
		result += string(rune(hashed))
	}
	return result

}
