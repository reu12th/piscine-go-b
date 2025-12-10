package piscine

func TrimAtoi(s string) int {
	sign := 1
	num := 0
	startNum := false

	for _, i := range s {
		if i == '-' && !startNum {
			sign = -1
		}

		if i == '0' && !startNum {
			continue
		}
		if i >= '0' && i <= '9' {
			startNum = true
			num = (num * 10) + int(i-'0')
		}
	}
	if !startNum {
		return 0
	}
	return (num * sign)
}
