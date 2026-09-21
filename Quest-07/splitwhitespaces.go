package piscine

func SplitWhiteSpaces(s string) []string {
	result := make([]string, 0)
	word := ""

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(s[i])
		}
	}

	if word != "" {
		result = append(result, word)
	}
	return result
}
