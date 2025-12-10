package piscine

func IsUpper(s string) bool {
	if len(s) <= 0 {
		return false
	}
	for _, a := range s {
		if a < 'A' || a > 'Z' {
			return false
		}
	}
	return true
}
