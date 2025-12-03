### 🟡 35% Tier: `findprevprime.go`

**Concept:** This function finds the largest prime number that is less than or equal to a given number `nb`. A prime number is a number greater than 1 that has no positive divisors other than 1 and itself.

**Line-by-Line Breakdown:**

1.  **`package piscine`**: Declares that this function is part of the `piscine` package.
2.  **`func FindPrevPrime(nb int) int {`**: Defines the function `FindPrevPrime`. It takes an integer `nb` as input and returns an integer (the prime found).
3.  **`for i := nb; i >= 2; i-- {`**: This is the main search loop.
    * **Initialization (`i := nb`)**: Instead of starting from 2 and counting up, we start at the input number `nb` and count *downwards*. This guarantees that the first prime number we find is the largest possible one less than or equal to `nb`.
    * **Condition (`i >= 2`)**: The smallest prime number is 2. If we go below 2, we stop searching.
    * **Decrement (`i--`)**: We decrease `i` by 1 in each iteration to check the next lower number.
4.  **`isPrime := true`**: We set a boolean flag `isPrime` to `true`. This assumes the current number `i` is a prime number until we find evidence to the contrary (a divisor).
5.  **`for j := 2; j < i; j++ {`**: This nested loop checks if `i` is divisible by any number `j`.
    * **Initialization (`j := 2`)**: We start dividing by 2 because every number is divisible by 1.
    * **Condition (`j < i`)**: We check all numbers less than `i` as potential divisors.
    * **Increment (`j++`)**: We check the next potential divisor (3, 4, 5, etc.).
6.  **`if i%j == 0 {`**: The modulo operator `%` returns the remainder of the division `i / j`. If the remainder is 0, it means `i` is perfectly divisible by `j`.
7.  **`isPrime = false`**: Since we found a divisor other than 1 and itself, `i` is **not** a prime number. We set the flag to `false`.
8.  **`break`**: We stop the inner loop immediately. We don't need to find *all* divisors; finding just one is enough to prove it's not prime. This optimization saves processing time.
9.  **`}`**: Closes the inner `if`.
10. **`}`**: Closes the inner `for` loop.
11. **`if isPrime {`**: After the inner loop finishes, we check the `isPrime` flag.
    * If it is still `true`, it means the inner loop completed without finding any divisors. Therefore, `i` is a prime number.
12. **`return i`**: We return `i` immediately because it is the prime number we were looking for.
13. **`}`**: Closes the outer `for` loop.
14. **`return 0`**: If the outer loop finishes (meaning `i` went below 2) without returning a prime, it implies no prime was found (e.g., if input was 1, 0, or negative). We return 0 as per the instructions.
