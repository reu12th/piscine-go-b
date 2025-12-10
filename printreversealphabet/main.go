package main

import "github.com/01-edu/z01"

func printReversalAlphabet() {
	for ch := 'z'; ch >= 'a'; ch-- {
		z01.PrintRune(ch)
	}
	z01.PrintRune(('\n'))
}

func main() {
	printReversalAlphabet()
}
