package piscine

func IsNumeric(s string) bool {
	if len(s) <= 0 {
		return false
	}
	for _, a := range s {
		if a < '0' || a > '9' {
			return false
		}
	}
	return true
}
