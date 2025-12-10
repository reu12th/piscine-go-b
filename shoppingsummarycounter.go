package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	item := make(map[string]int)

	var temp string

	for _, char := range str {
		if char == 32 {
			item[temp]++
			temp = ""
		} else if char != 32 {
			temp += string(byte(char))
		}
	}
	item[temp]++

	return item
}
