package piscine

func Compact(ptr *[]string) int {
	count := 0
	var args []string

	for _, char := range *ptr {
		if char != "" {
			args = append(args, char)
			count++
		}
		*ptr = args
	}
	return count
}
