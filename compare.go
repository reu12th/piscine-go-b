package piscine

func Compare(a, b string) int {
	// Set len a to minimum length
	minLen := len(a)

	// compare lengths with len of string b, to get string with smallest length
	if len(b) < len(a) {
		minLen = len(b)
	}

	for i := 0; i < minLen; i++ {
		if a[i] != b[i] {
			if a[i] < b[i] {
				return -1
			} else {
				return 1
			}
		}
	}

	if len(a) == len(b) {
		return 0
	} else if len(a) < len(b) {
		return -1
	} else {
		return 1
	}
}
