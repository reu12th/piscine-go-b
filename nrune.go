package piscine

func NRune(s string, n int) rune {
	if n < 0 {
		return 0
	}
	for i, a := range s {
		if i == n-1 {
			return a
		}
	}
	return 0
}
