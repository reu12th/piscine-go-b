package piscine

func BasicAtoi2(s string) int {
	var n int
	for _, digit := range s {
		if digit > '9' || digit < '0' {
			return 0
		}
		n = n*10 + int(digit-'0')
	}
	return n
}
