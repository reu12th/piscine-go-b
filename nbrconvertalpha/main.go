package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	paramsList := os.Args[1:]
	allCaps := false

	if len(paramsList) > 0 && paramsList[0] == "--upper" {
		allCaps = true
		paramsList = paramsList[1:]
	}

	for _, j := range paramsList {
		val, isValidInt := isInt(j)
		if isValidInt && (val >= 1 && val <= 26) {
			if val == 0 {
				z01.PrintRune('0')
				continue
			}
			if allCaps {
				z01.PrintRune(rune(val + 64))
			} else {
				z01.PrintRune(rune(val + 96))
			}
		} else {
			z01.PrintRune(' ')
		}
	}
	if len(paramsList) > 0 {
		z01.PrintRune('\n')
	}
}

// Check if a string if a stringed form of an integer
// If it is, convert the string to an integer, and return the int value of the string and boolean value of true
// If not, return 0, and false

func isInt(s string) (int, bool) {
	strLen := len(s)
	if strLen == 0 {
		return 0, false
	}

	result := 0

	for i := 0; i < strLen; i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		num := int((s[i]) - '0')
		result = result*10 + num
	}
	return result, true
}
