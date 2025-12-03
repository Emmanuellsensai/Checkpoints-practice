package main

/*
PSEUDOCODE FOR main FUNCTION:

STEP 1 — Call the Itoa function with test numbers.
STEP 2 — Print results to check correctness.
*/

import "fmt"

func main() {

	// STEP: Test various values
	fmt.Println(Itoa(12345))     // Expected: "12345"
	fmt.Println(Itoa(0))         // Expected: "0"
	fmt.Println(Itoa(-1234))     // Expected: "-1234"
	fmt.Println(Itoa(987654321)) // Expected: "987654321"
}
