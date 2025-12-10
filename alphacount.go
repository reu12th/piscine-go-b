package piscine

func AlphaCount(s string) int {
	counter := 0

	for _, a := range s {
		if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') {
			counter++
		}
	}
	return counter
}
