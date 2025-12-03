package piscine

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	result := ""
	for i, r := range s {
		// Check if the character is Uppercase
		if r >= 'A' && r <= 'Z' {
			// If it's not the first letter, and the previous letter was not uppercase
			// (OR depending on subject, simply if the previous was lowercase)
			if i != 0 && (s[i-1] >= 'a' && s[i-1] <= 'z') {
				result += "_"
			}
			// Append the Lowercase version of the character
			result += string(r + 32)
		} else {
			// Append character as is
			result += string(r)
		}
	}
	return result
}
