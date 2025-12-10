package piscine

func ConcatParams(args []string) string {
	newStr := ""

	for i := range args {
		if i > 0 {
			newStr += "\n"
		}
		newStr += args[i]
	}
	return newStr
}
