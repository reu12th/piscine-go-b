package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	num := toDecimal(nbr, baseFrom)

	return fromDecimal(num, baseTo)
}

func toDecimal(nbr string, base string) int {
	baseLen := len(base)
	valueMap := make(map[rune]int)

	for i, r := range base {
		valueMap[r] = i
	}

	result := 0
	for _, r := range nbr {
		result = result*baseLen + valueMap[r]
	}
	return result
}

func fromDecimal(n int, base string) string {
	if n == 0 {
		return string(base[0])
	}

	baseLen := len(base)
	result := ""

	for n > 0 {
		result = string(base[n%baseLen]) + result
		n /= baseLen
	}

	return result
}
