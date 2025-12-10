package piscine

func SplitWhiteSpaces(s string) []string {
	var stringSlice []string
	wordIndex := -1

	for i, ch := range s {
		if ch != ' ' && ch != '\t' && ch != '\n' {
			if wordIndex == -1 {
				wordIndex = i
			}
		} else {
			if wordIndex != -1 {
				stringSlice = append(stringSlice, s[wordIndex:i])
				wordIndex = -1
			}
		}
	}

	if wordIndex != -1 {
		stringSlice = append(stringSlice, s[wordIndex:])
	}

	return stringSlice
}
