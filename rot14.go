package piscine

func Rot14(s string) string {
	result := ""

	for _, char := range s {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			char += 14
			if char > 'z' {
				char -= 26
			}
		} else if char >= 'A' && char <= 'Z' {
			char += 14
			if char > 'Z' {
				char -= 26
			}
		}

		result += string(char)
	}

	return result
}
