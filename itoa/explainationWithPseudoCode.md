
---

# **🧠 Itoa Function — Full Explanation**

This document explains **how our custom `Itoa` function works**, following a clear, beginner-friendly, and Piscine-approved teaching style.

---

## **📌 What Is Itoa?**

`Itoa` means **Integer to ASCII**, or more simply:

> Convert an `int` → into a `string`.

We rebuild it *manually* without using Go’s built-in conversion functions.

---

## **📘 High-Level Steps**

1. Handle the special case: **n = 0**
2. Check if the number is **negative**
3. Extract digits from **last to first**
4. Convert each digit to a **character**
5. Append digits in **reverse order**
6. Add a **minus sign** if needed
7. Reverse the final slice of characters
8. Convert slice → **string**

---

# **🧩 Detailed Pseudocode Explanation**

## **1️⃣ Step 1 — Special Case: n == 0**

We return `"0"` immediately, because the loop logic will not work for zero.

```go
if n == 0 {
    return "0"
}
```

---

## **2️⃣ Step 2 — Check If Number Is Negative**

If the number is negative:

* Mark a flag (`isNegative = true`)
* Convert it to positive for easier digit extraction

```go
isNegative := false
if n < 0 {
    isNegative = true
    n = -n
}
```

---

## **3️⃣ Step 3 — Extract Digits One-by-One**

We repeatedly take:

* `n % 10` → gets **last digit**
* `'0' + digit` → converts digit into a **character**
* Append characters in **reverse order**
* Remove last digit using `n = n / 10`

```go
digits := []rune{}
for n > 0 {
    lastDigit := n % 10
    char := rune('0' + lastDigit)
    digits = append(digits, char)
    n = n / 10
}
```

**Example:**
Input: `123`
Extracted order: `'3'`, `'2'`, `'1'`
We will reverse later.

---

## **4️⃣ Step 4 — Add '-' If Number Was Negative**

We append the minus sign *before reversing*, so it also gets correctly positioned.

```go
if isNegative {
    digits = append(digits, '-')
}
```

---

## **5️⃣ Step 5 — Reverse the Characters**

The digits are stored backward, so we reverse the whole slice.

```go
for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
    digits[i], digits[j] = digits[j], digits[i]
}
```

---

## **6️⃣ Step 6 — Convert Digits to a Final String**

The last step: return the completed string.

```go
return string(digits)
```

---

# **✔ Final Code (For Reference)**

```go
package piscine

/*
Step 1: Handle the special case where n == 0
*/
func Itoa(n int) string {
    if n == 0 {
        return "0"
    }

    /*
    Step 2: Detect if the number is negative
    */
    isNegative := false
    if n < 0 {
        isNegative = true
        n = -n
    }

    /*
    Step 3: Extract digits from last to first
    */
    digits := []rune{}
    for n > 0 {
        lastDigit := n % 10
        char := rune('0' + lastDigit)
        digits = append(digits, char)
        n = n / 10
    }

    /*
    Step 4: Append '-' if number was negative
    */
    if isNegative {
        digits = append(digits, '-')
    }

    /*
    Step 5: Reverse the slice
    */
    for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
        digits[i], digits[j] = digits[j], digits[i]
    }

    /*
    Step 6: Convert slice to string
    */
    return string(digits)
}
```

---

# **🎉 Teaching Summary**

Emphasize:

### **✔ The algorithm works WITHOUT using strconv**

### **✔ We manually extract digits with % and /**

### **✔ Strings can be built from runes**

### **✔ Negative numbers need special handling**

### **✔ Reversal is required because digits come out backwards**

---

