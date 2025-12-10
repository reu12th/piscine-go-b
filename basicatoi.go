package piscine

func BasicAtoi(s string) int {
	var n int
	for _, digit := range s {
		n = n*10 + int(digit-'0')
	}
	return n
}
