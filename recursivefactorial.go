package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	if nb == 0 {
		return 1
	}

	maxInt := int(^uint(0) >> 1)

	// Factorials larger than 20 overflow on a 64-bit int.
	if nb > 20 {
		return 0
	}

	previous := RecursiveFactorial(nb - 1)

	if previous > maxInt/nb {
		return 0
	}

	return nb * previous
}
