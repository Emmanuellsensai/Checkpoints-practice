package main

/*
PSEUDOCODE FOR Itoa FUNCTION

STEP 1 — If the number is 0:
          Return "0" immediately.

STEP 2 — Check if the number is negative.
          If yes:
             - Record that it is negative
             - Convert n to positive for easier processing

STEP 3 — Extract digits one-by-one by using:
             digit = n % 10
             n = n / 10

STEP 4 — Convert the extracted digit to its character equivalent:
             char = '0' + digit

STEP 5 — Append each converted digit to a temporary string
          (Note: digits come out REVERSED)

STEP 6 — If the original number was negative:
             Add '-' sign

STEP 7 — Reverse the built string because digits were collected backwards.

STEP 8 — Return the final string.
*/

func Itoa(n int) string {

    // Handle zero case early
    if n == 0 {
        return "0"
    }

    negative := false

    // STEP: Handle negative numbers
    if n < 0 {
        negative = true
        n = -n // Work with positive version
    }

    result := ""

    // STEP: Extract digits and build result in reverse
    for n > 0 {
        digit := n % 10
        result += string(rune('0' + digit))
        n /= 10
    }

    // STEP: Add negative sign if needed
    if negative {
        result += "-"
    }

    // STEP: Reverse the string because digits were appended backwards
    reversed := ""
    for i := len(result) - 1; i >= 0; i-- {
        reversed += string(result[i])
    }

    return reversed
}
