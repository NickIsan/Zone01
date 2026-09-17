package piscine

func Split(s, sep string) []string {
	result := make([]string, 0)

	word := ""

	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			result = append(result, word)
			word = ""
			i += len(sep) - 1
		} else {
			word += string(s[i])
		}
	}
	result = append(result, word)

	return result
}
