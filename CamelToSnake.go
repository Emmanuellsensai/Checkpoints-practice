package piscine

func CamelToSnakeCase(s string) string {
	// 1. Empty Check
	if s == "" {
		return ""
	}

	// --- VALIDATION PHASE ---
	// Before converting, we check if the string adheres to strict CamelCase rules.
	// If it violates any rule, we return the original string 's' unchanged.

	for i := 0; i < len(s); i++ {
		// Rule A: Numbers or punctuation are not allowed.
		// We check if the character is OUTSIDE the range of A-Z AND a-z.
		if (s[i] < 'A' || s[i] > 'Z') && (s[i] < 'a' || s[i] > 'z') {
			return s
		}

		// Checks specific to Uppercase letters
		if s[i] >= 'A' && s[i] <= 'Z' {
			// Rule B: The string cannot END with a capitalized letter.
			// If we are at the last index (len-1) and it's uppercase -> Invalid.
			if i == len(s)-1 {
				return s
			}

			// Rule C: No two capitalized letters can follow each other directly.
			// We check the next character (s[i+1]).
			// Note: This access is safe because we already verified (i != len-1) above.
			if s[i+1] >= 'A' && s[i+1] <= 'Z' {
				return s
			}
		}
	}

	// --- CONVERSION PHASE ---
	// If the string survived validation, we construct the snake_case version.

	result := ""
	for i, r := range s {
		// If character is Uppercase
		if r >= 'A' && r <= 'Z' {
			// 1. Insert Underscore
			// We only insert if it's NOT the very first letter.
			if i != 0 {
				result += "_"
			}
			// 2. Convert to Lowercase
			// 'A'(65) + 32 = 'a'(97). We shift the ASCII value.
			result += string(r + 32)
		} else {
			// If character is Lowercase, append it exactly as is.
			result += string(r)
		}
	}

	return result
}
