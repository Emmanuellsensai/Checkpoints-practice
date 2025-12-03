package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	// --- 1. VALIDATION STEP ---
	for i := 0; i < len(s); i++ {
		// Rule A: Numbers or punctuation are not allowed (Must be A-Z or a-z)
		if (s[i] < 'A' || s[i] > 'Z') && (s[i] < 'a' || s[i] > 'z') {
			return s
		}

		// Checks specific to Uppercase letters
		if s[i] >= 'A' && s[i] <= 'Z' {
			// Rule B: The word does not end on a capitalized letter
			if i == len(s)-1 {
				return s
			}
			// Rule C: No two capitalized letters shall follow directly each other
			// We check s[i+1], which is safe because we already checked (i == len-1) above
			if s[i+1] >= 'A' && s[i+1] <= 'Z' {
				return s
			}
		}
	}

	// --- 2. CONVERSION STEP ---
	result := ""
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			// If it is uppercase:
			// 1. Add underscore (only if it's not the very first letter)
			if i != 0 {
				result += "_"
			}
			// 2. Convert to lowercase and append
			result += string(r + 32)
		} else {
			// If it is lowercase, append as is
			result += string(r)
		}
	}
	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
