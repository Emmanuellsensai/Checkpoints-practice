# 🔍 Extensive Explanation For FindPrevPrime

**Concept:**
A **Prime Number** is a number greater than 1 that has no divisors other than 1 and itself. The goal of `FindPrevPrime` is to find the largest prime number that is less than or equal to the input `nb`.



[Image of prime numbers on number line]


**Line-by-Line Breakdown:**

1.  **`package piscine`**: This line declares that the function belongs to the `piscine` package, which is standard for these exercises.

2.  **`func FindPrevPrime(nb int) int {`**: Defines the function signature. It takes an integer `nb` as input and returns an integer (the prime found).

3.  **`if nb < 2 { return 0 }`**: This handles invalid inputs.
    * By definition, the smallest prime number is 2.
    * If the input is `1`, `0`, or `-5`, no prime exists below it. The instructions typically require returning `0` in this case.

4.  **`for i := nb; i >= 2; i-- {`**: This is the main search loop.
    * **Initialization (`i := nb`)**: We start checking from the number given. If `nb` itself is prime, we return it.
    * **Condition (`i >= 2`)**: We stop if we go below 2.
    * **Decrement (`i--`)**: We count backwards (`nb`, `nb-1`, `nb-2`...) because we want the *previous* prime (closest to `nb`).

5.  **`isPrime := true`**: This boolean flag acts as our "innocent until proven guilty" marker. We assume the current number `i` is prime. We will try to disprove this in the next loop.

6.  **`for j := 2; j*j <= i; j++ {`**: This inner loop checks for factors.
    * **`j := 2`**: We start dividing by 2 (dividing by 1 is useless as everything is divisible by 1).
    * **`j*j <= i`**: This is a critical optimization. Instead of checking all numbers up to `i`, we stop at the square root of `i`.
        * *Why?* If `i` is `36`, its factors are `(2,18)`, `(3,12)`, `(4,9)`, `(6,6)`. Notice that after the square root (6), the factors just repeat in reverse order. If we haven't found a factor by the square root, we never will.
    * **`j++`**: We check 2, then 3, then 4, etc.

7.  **`if i%j == 0 {`**: The Modulo operator `%` gives the remainder of division.
    * If `i % j == 0`, it means `i` divides perfectly by `j`.
    * Example: `9 % 3 == 0`. Since 9 has a divisor (3) other than 1 and 9, it is **not** prime.

8.  **`isPrime = false; break`**:
    * We mark the number as "not prime".
    * `break` exits the *inner* loop immediately. We don't need to find *all* divisors; finding just one is enough to prove it's composite. This saves time.

9.  **`if isPrime { return i }`**:
    * After the inner loop finishes, we check our flag.
    * If `isPrime` is still `true`, it means the inner loop finished without finding any divisors.
    * Therefore, `i` is prime. We return it immediately as the answer.

10. **`return 0`**: This line is technically unreachable if `nb >= 2` (since 2 is prime and the loop goes down to 2), but it satisfies the compiler's requirement for a return value.
