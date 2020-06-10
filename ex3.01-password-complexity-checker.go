package main

import (
	"fmt"
	"unicode"
)

/*
strong password requirements:-
- Has a lowercase letter
- Has an uppercase letter
- Has a number
- Has a symbol
- Be 8 or more characters long
*/

func passwordChecker(pw string) bool {
	// We need to convert password string into `rune`, which is a type that is safe for multi-byte (UTF-8) characters
	// pwR = password rune
	pwR := []rune(pw)
	if len(pwR) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, v := range pwR {
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}
		// we will also accept punctuation marks for symbols
		if unicode.IsSymbol(v) || unicode.IsPunct(v) {
			hasSymbol = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	// passwordChecker(" ")
	if passwordChecker("ThisI$@passw0rd!") {
		fmt.Println("Strong password!")
	} else {
		fmt.Println("Weak password!")
	}
}
