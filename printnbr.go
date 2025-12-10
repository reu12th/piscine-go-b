package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
	}
	PrintDigits(n)
}

func PrintDigits(n int) {
	if n/10 != 0 {
		PrintDigits(n / 10)
	}
	remainder := n % 10
	if n < 0 {
		remainder = -remainder
	}
	z01.PrintRune(rune('0' + remainder))
}
