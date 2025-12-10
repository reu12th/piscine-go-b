package piscine

func Rot14(s string) string {
	char := []byte(s)
	for i := 0; i < len(char); i++ {
		if char[i] >= 'a' && char[i] <= 'l' || char[i] >= 'A' && char[i] <= 'L' {
			char[i] += 14
		} else if char[i] > 'l' && char[i] <= 'z' || char[i] > 'L' && char[i] <= 'Z' {
			char[i] -= 12
		}
	}
	return string(char)
}
