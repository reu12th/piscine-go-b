package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	comb := make([]int, n)
	isFirst := true

	PrintoutCombN(comb, 0, 0, n, &isFirst)
	z01.PrintRune('\n')
}

func PrintoutCombN(comb []int, pos, start, n int, isFirst *bool) {
	if pos == n {
		if !*isFirst {
			z01.PrintRune((','))
			z01.PrintRune((' '))
		}
		*isFirst = false
		for _, digit := range comb {
			z01.PrintRune(rune(digit + '0'))
		}
		return
	}

	for i := start; i <= 10-(n-pos); i++ {
		comb[pos] = i
		PrintoutCombN(comb, pos+1, i+1, n, isFirst)
	}
}
