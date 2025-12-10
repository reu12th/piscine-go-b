package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n <= 0 {
		z01.PrintRune('0')
	}

	numbers := []int{}

	for n > 0 {
		num := n % 10
		numbers = append(numbers, num)
		n /= 10
	}
	SortIntegerTable(numbers)
	for _, i := range numbers {
		z01.PrintRune(rune(i) + '0')
	}
}
