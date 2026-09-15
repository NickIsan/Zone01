package piscine

func Index(s string, toFind string) int {
	sRunes := []rune(s)
	findRunes := []rune(toFind)

	if len(findRunes) == 0 {
		return 0
	}

	// Does toFind fit inside s? From which index will I start comparing?
	for i := 0; i+len(findRunes) <= len(sRunes); i++ {
		found := true

		// Are all the characters the same?
		for j := 0; j < len(findRunes); j++ {
			if sRunes[i+j] != findRunes[j] {
				found = false
				break
				// Stops the for j
			}
		}

		if found {
			return i
		}
	}

	return -1
}
