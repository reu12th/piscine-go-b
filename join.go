package piscine

func Join(strs []string, sep string) string {
	joined := ""

	for i, char := range strs {
		if len(strs)-1 == i {
			joined += char
		} else {
			joined += char + sep
		}
	}
	return joined
}
