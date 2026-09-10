package piscine

func StrRev(s string) string {
	runes := []rune(s)
	result := ""

	for i := len(runes) - 1; i >= 0; i-- {
		result += string(runes[i])
	}

	return result
}
