package piscine

func IsAlpha(s string) bool {
	if len(s) <= 0 {
		return false
	}
	for _, a := range s {
		if (a < 'a' || a > 'z') && (a < 'A' || a > 'Z') && (a < '0' || a > '9') {
			return false
		}
	}
	return true
}
