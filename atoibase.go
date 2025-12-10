package piscine

func AtoiBase(s string, base string) int {

	if !isValidBase(base) {
		return 0
	}

	baseRunes := []rune(base)
	baseLen := len(baseRunes)

	// Convert each character to its numeric value
	result := 0
	for _, ch := range s {
		value := indexInBase(ch, baseRunes)
		if value == -1 { // character not in base
			return 0
		}
		result = result*baseLen + value
	}

	return result
}

// Check base validity
func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, ch := range base {
		if ch == '+' || ch == '-' || seen[ch] {
			return false
		}
		seen[ch] = true
	}
	return true
}

// Return index of a rune in base, or -1 if not found
func indexInBase(ch rune, base []rune) int {
	for i, b := range base {
		if ch == b {
			return i
		}
	}
	return -1
}
