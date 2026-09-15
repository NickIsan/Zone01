package piscine

func TrimAtoi(s string) int {
	result := 0
	sign := 1
	foundDigit := false

	for _, r := range s {
		if r == '-' && !foundDigit {
			sign = -1
			continue
		}

		if r == '+' && !foundDigit {
			sign = 1
			continue
		}

		if r >= '0' && r <= '9' {
			foundDigit = true
			result = result*10 + int(r-'0')
		}
	}

	return result * sign
}
