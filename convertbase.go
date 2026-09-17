package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Convert nbr to decimal
	number := 0

	for _, char := range nbr {
		for i, baseChar := range baseFrom {
			if char == baseChar {
				number = number*len(baseFrom) + i
				break
			}
		}
	}

	// Convert decimal to baseTo
	if number == 0 {
		return string(baseTo[0])
	}

	result := ""

	for number > 0 {
		remainder := number % len(baseTo)
		result = string(baseTo[remainder]) + result
		number = number / len(baseTo)
	}

	return result
}
