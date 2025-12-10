package piscine

func ToLower(s string) string {
	newStr := []rune{}
	for _, a := range s {
		if a >= 'A' && a <= 'Z' {
			a += 32
		}
		newStr = append(newStr, a)
	}
	return string(newStr)
}
