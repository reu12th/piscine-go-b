package piscine

func JumpOver(str string) string {
	if len(str) == 0 || len(str) < 3 {
		return "\n"
	}

	thirdChar := ""
	for i := 2; i < len(str); i += 3 {
		if str[i] != ' ' {
			thirdChar += string(str[i])
		}
	}
	return thirdChar + "\n"
}
