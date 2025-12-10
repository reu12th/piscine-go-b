package piscine

func Capitalize(s string) string {
	strRune := []rune(s)
	isInWord := false

	for i, a := range strRune {
		if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9') {
			if !isInWord {
				if a >= 'a' && a <= 'z' {
					strRune[i] = a - 32
				}
				isInWord = true
			} else {
				if a >= 'A' && a <= 'Z' {
					strRune[i] = a + 32
				}
			}
		} else {
			isInWord = false
		}
	}
	return (string(strRune))
}
