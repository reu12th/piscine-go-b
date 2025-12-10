package piscine

func IsPrintable(s string) bool {
	for _, a := range s {
		if a <= rune(31) || a == rune(127) {
			return false
		}
	}
	return true
}
