package piscine

func Atoi(s string) int {
	runes := []rune(s)
	if len(runes) == 0 {
		return 0
	}

	sign := 1
	start := 0

	if runes[0] == '-' {
		sign = -1
		start = 1
	} else if runes[0] == '+' {
		start = 1
	}

	result := 0
	for i := start; i < len(runes); i++ {
		if runes[i] < '0' || runes[i] > '9' {
			return 0
		}
		result = result*10 + int(runes[i]-'0')
	}

	return sign * result
}
