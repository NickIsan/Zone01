package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 0
	}

	if nb%2 == 0 {
		nb++
	}

	for {
		if IsPrime(nb) {
			return nb
		}

		nb += 2
	}
}
