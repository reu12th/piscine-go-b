package piscine

func ReverseMenuIndex(menu []string) []string {
	arg := len(menu) - 1

	var arr []string
	for i := arg; i >= 0; i-- {
		arr = append(arr, menu[i])
	}
	return arr
}
