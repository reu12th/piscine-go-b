package piscine

func Abort(a, b, c, d, e int) int {
	entries := []int{a, b, c, d, e}

	SortIntegerTable(entries)

	median := entries[2]

	return median
}
