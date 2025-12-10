package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 || nb == 3 {
		return true
	}
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}

	// Check if numbers greater than 3, that are not divisible by 2 or 3 are primes

	for i := 5; i*i <= nb; i += 6 {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
	}
	return true
}
