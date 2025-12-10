package piscine

func IterativeFactorial(nb int) int {
	factorial := 1
	if nb < 0 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		previous := factorial

		factorial = factorial * i
		if factorial/i != previous {
			return 0
		}
	}
	return factorial
}
