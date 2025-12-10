package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}

	left, right := 1, nb/2+1
	return SolveSqrt(left, right, nb)
}

func SolveSqrt(left, right, x int) int {
	if left > right {
		return 0
	}
	mid := (left + right) / 2
	if mid*mid == x {
		return mid
	}
	if mid*mid < x {
		left = mid + 1
	} else {
		right = mid - 1
	}
	return SolveSqrt(left, right, x)
}
