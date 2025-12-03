Here is the extensive, line-by-line explanation of the `CamelToSnakeCase` function found in the Canvas `cameltosnakecase.go`.

### 🟡 35% Tier: `cameltosnakecase.go`

**Concept:** This function takes a string written in "CamelCase" format and converts it into "snake_case". This involves two main actions:
1.  Inserting underscores `_` before new words (identified by uppercase letters).
2.  Converting those uppercase letters to lowercase.

**Line-by-Line Breakdown:**

1.  **`package piscine`**: Defines the package name. This function belongs to the `piscine` library.
2.  **`func CamelToSnakeCase(s string) string {`**: Defines the function. It accepts one string `s` as input and returns a string.
3.  **`if s == "" { return "" }`**: Safety check. If the input string is empty, there is nothing to convert. We return an empty string immediately to avoid index errors later.
4.  **`result := ""`**: Initializes an empty string variable `result`. We will build the new "snake_case" version of the string here, character by character.
5.  **`for i, r := range s {`**: Starts a loop iterating through the input string `s`.
    * `i` represents the current **index** (0, 1, 2...).
    * `r` represents the current **rune** (character).
6.  **`if r >= 'A' && r <= 'Z' {`**: This condition checks if the current character `r` is an Uppercase letter (between 'A' and 'Z'). This is our primary trigger for finding the start of a new word in CamelCase.
7.  **`if i != 0 && (s[i-1] >= 'a' && s[i-1] <= 'z') {`**: This nested condition determines if we need to insert an underscore.
    * **`i != 0`**: We never insert an underscore at the very beginning of the string (e.g., "Camel" becomes "camel", not "_camel").
    * **`s[i-1] >= 'a' && s[i-1] <= 'z'`**: We check the **previous** character (`s[i-1]`). We only add an underscore if the previous character was a **lowercase letter**.
        * This logic ensures that transitions like `camelCase` (l -> C) get an underscore (`camel_case`).
        * It prevents underscores inside acronyms if they are typed consecutively (depending on specific rules, but strictly this logic looks for `lower` -> `Upper` transitions).
8.  **`result += "_"`**: If the condition above is true (we found a word boundary), we append the underscore separator to `result`.
9.  **`result += string(r + 32)`**: This handles the case conversion.
    * In ASCII, Uppercase letters are 32 values lower than Lowercase letters (e.g., 'A' is 65, 'a' is 97).
    * By adding `32` to the rune `r`, we mathematically convert it to lowercase.
    * We wrap it in `string()` to append it to our result.
10. **`} else {`**: If the current character `r` was **not** Uppercase (it is lowercase, a number, or a symbol).
11. **`result += string(r)`**: We simply append the character to `result` exactly as it is, without changing case or adding underscores.
12. **`}`**: Closes the loop.
13. **`return result`**: Returns the fully constructed snake_case string.
