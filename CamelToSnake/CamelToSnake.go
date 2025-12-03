package piscine

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	result := ""
	for i, r := range s {
		// If Uppercase...
		if r >= 'A' && r <= 'Z' {
			// AND not at start, AND previous char wasn't underscore...
			// AND (typically) we check if the char is actually the start of a new word
			// (Handling weird cases like CamelCase vs CamelCASE is tricky, standard logic below)
			if i != 0 && (s[i-1] < 'A' || s[i-1] > 'Z') {
				result += "_"
			}
		}
		result += string(r)
	}
	return result
}
