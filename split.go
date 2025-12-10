package piscine

func Split(s, sep string) []string {
	sepLen := len(sep)
	str := []string{}
	counter := 0

	for i := range s {
		if i+sepLen < len(s) {
			if s[i:i+sepLen] == sep {
				counter++
			}
		}
	}

	str = make([]string, counter+1)
	strIndex := 0
	wordIndex := 0

	for i := range s {
		if i+sepLen < len(s) {
			if s[i:i+sepLen] == sep {
				str[strIndex] = s[wordIndex:i]
				strIndex++
				i += sepLen
				wordIndex = i
			}
		}
	}
	str[strIndex] = s[wordIndex:]

	return str
}
