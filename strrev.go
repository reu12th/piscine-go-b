package piscine

func StrRev(s string) string {
	runeS := []rune(s)
	revers := make([]rune, len(runeS))
	for i := 0; i < len(runeS); i++ {
		revers[i] = runeS[len(runeS)-i-1]
	}
	return string(revers)
}
