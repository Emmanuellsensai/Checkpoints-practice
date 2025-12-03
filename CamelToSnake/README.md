## 🟡 35% Tier: cameltosnakecase.go

Concept: This function transforms a string formatted in "CamelCase" (capitalized words joined together) into "snake_case" (lowercase words separated by underscores).

package piscine: Declares that this code is part of the piscine package, allowing it to be exported and used by other files in the same package.

func CamelToSnakeCase(s string) string {: Defines a function named CamelToSnakeCase that accepts one argument s (a string) and returns a string.

if s == "" { return "" }: This is a standard "edge case" check. If the input string is empty, there is nothing to convert, so it immediately returns an empty string to avoid errors later.

result := "": Initializes an empty string variable named result. This will act as a buffer where we build the new snake_case string character by character.

for i, r := range s {: Starts a loop that iterates through every character in the input string s.

i is the index (position) of the current character (0, 1, 2...).

r is the "rune" (the character itself).

if r >= 'A' && r <= 'Z' {: Checks if the current character r is an uppercase letter. In ASCII/Unicode, capital 'A' through 'Z' are sequential, so this range check identifies them.

if i != 0 && (s[i-1] < 'A' || s[i-1] > 'Z') {: This complex condition decides if an underscore needs to be inserted before the current uppercase letter.

i != 0: We never put an underscore at the very beginning of the string (e.g., "Camel" stays "Camel", it doesn't become "_Camel").

s[i-1] < 'A' || s[i-1] > 'Z': This checks the previous character. If the character before the current one was not uppercase (meaning it was lowercase or a number), it implies a word boundary. For example, in "camelCase", when we hit 'C', the previous char is 'l' (lowercase). This triggers the insertion.

Note: This specific logic skips adding underscores between consecutive capitals (like "HTML"), treating them as one word block depending on specific exercise variants.

result += "_": If the condition above is met, append an underscore to the result string.

}: Closes the inner if.

}: Closes the outer if.

result += string(r): Appends the current character r to the result string. We cast r (which is a rune/int32) back to a string type to add it.

}: Closes the for loop.

return result: Returns the final constructed string.
