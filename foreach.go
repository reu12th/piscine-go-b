package piscine

func ForEach(f func(int), arr []int) {
	for _, r := range arr {
		f(r)
	}
}
