package piscine

func IsLower(s string) bool {
	if len(s) <= 0 {
		return false
	}
	for _, a := range s {
		if a < 'a' || a > 'z' {
			return false
		}
	}
	return true
}
