package piscine

func BasicJoin(elems []string) string {
	newStr := ""
	for _, r := range elems {
		newStr += r
	}
	return newStr
}
