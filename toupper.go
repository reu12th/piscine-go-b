package piscine

func ToUpper(s string) string {
	newStr := []rune{}

	for _, a := range s {
		if a >= 'a' && a <= 'z' {
			a -= 32
		}
		newStr = append(newStr, a)
	}
	return string(newStr)
}
