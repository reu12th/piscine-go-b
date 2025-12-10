package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	if nb%2 == 0 {
		nb++
	}
	if IsPrime(nb) {
		return nb
	}
	return FindNextPrime(nb + 2)
}
