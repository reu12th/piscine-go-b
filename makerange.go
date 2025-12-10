package piscine

func MakeRange(min, max int) []int {
	arrSize := max - min
	result := []int(nil)

	if min < max {

		result = make([]int, arrSize)
		val := min
		for i := 0; i < arrSize; i++ {
			result[i] = val
			val++
		}
	}
	return result
}
