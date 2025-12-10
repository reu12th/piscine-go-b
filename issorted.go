package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	count := 1
	count2 := 1
	count3 := 1
	for k, r := range tab {
		if k != len(tab)-1 {
			if f(r, tab[k+1]) < 0 {
				count++
			}
			if f(r, tab[k+1]) > 0 {
				count2++
			}
			if f(r, tab[k+1]) == 0 {
				count3++
			}
		}
	}
	return count == len(tab) || count2 == len(tab) || count3 == len(tab)
}
