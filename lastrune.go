package piscine

func LastRune(s string) rune {
	lastRune := rune(0)

	for _, a := range s {
		lastRune = a
	}
	return lastRune
}
