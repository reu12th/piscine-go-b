package piscine

func Index(s string, toFind string) int {
	j := len(s)
	k := len(toFind)

	if j == 0 {
		return 0
	}
	if k > j {
		return -1
	}

	for i := 0; i < j-k; i++ {
		if s[i:i+k] == toFind {
			return i
		}
	}
	return -1
}
