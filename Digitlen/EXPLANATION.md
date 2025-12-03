### 🔍 Extensive Explanation

**Concept:**
The function `DigitLen` calculates how many digits are required to represent a number `n` in a specific numerical `base`. This is mathematically equivalent to finding how many times you can divide the number by the base before it becomes 0.

**Line-by-Line Breakdown:**

1.  **`package piscine`**: Declares that this function belongs to the `piscine` package.

2.  **`func DigitLen(n, base int) int {`**: Defines the function taking two integers and returning an integer (the length).

3.  **`if base < 2 || base > 36 { return -1 }`**:
    * This is a validity check. Numerical bases generally start at 2 (Binary) and go up to 36 (digits 0-9 plus letters A-Z).
    * If the base is invalid (like 1, 0, -5, or 100), the function returns `-1` to indicate an error.

4.  **`if n < 0 { n = -n }`**:
    * The length of a number's representation is usually concerned with its magnitude (digits), not the sign.
    * Example: Length of `-100` in base 10 is 3 (digits `1`, `0`, `0`).
    * We convert `n` to positive so our calculation loop works correctly.

5.  **`if n == 0 { return 1 }`**:
    * This handles the specific case of the number `0`.
    * If we skipped this and went straight to the loop `for n > 0`, the loop condition would fail immediately (since 0 is not > 0), and `count` would be 0.
    * However, the number `0` is written as "0", which has a length of **1 digit**. So we return 1 explicitly.

6.  **`count := 0`**: Initializes the counter.

7.  **`for n > 0 {`**: This is the core logic loop. It runs as long as `n` is positive.

8.  **`n /= base`**:
    * This uses integer division to chop off the last digit.
    * Example: `100` (base 10).
        * `100 / 10 = 10`.
        * `10 / 10 = 1`.
        * `1 / 10 = 0`.
    * It simulates shifting the number to the right in the given base.

9.  **`count++`**: Every time we successfully divide, it represents one digit that existed. We increment the count.

10. **`return count`**: Returns the total number of divisions performed, which equals the number of digits.
