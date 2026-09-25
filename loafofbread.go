package piscine

func LoafOfBread(str string) string {
	runes := []rune(str)
	if len(runes) == 0 {
		return "\n"
	}
	if len(runes) < 5 {
		return "Invalid Output\n"
	}

	var result []rune
	var block []rune
	skipNext := false

	i := 0
	for i < len(runes) {
		ch := runes[i]
		i++

		if skipNext {
			skipNext = false
			continue
		}

		if ch == ' ' {
			continue
		}

		block = append(block, ch)

		if len(block) == 5 {
			if len(result) > 0 {
				result = append(result, ' ')
			}
			result = append(result, block...)
			block = nil
			skipNext = true
		}
	}

	if len(block) > 0 {
		if len(result) > 0 {
			result = append(result, ' ')
		}
		result = append(result, block...)
	}

	result = append(result, '\n')
	return string(result)
}

// package piscine

// func LoafOfBread(str string) string {
// 	if str == "" {
// 		return "\n"
// 	}
// 	if len(str) < 5 {
// 		return "Invalid Output\n"
// 	}

// 	var result []byte
// 	var block []byte
// 	skipNext := false

// 	for i := 0; i < len(str); i++ {
// 		character := str[i]

// 		if skipNext == true {
// 			skipNext = false
// 			continue
// 		}

// 		if character == ' ' {
// 			continue
// 		}

// 		block = append(block, character)

// 		if len(block) == 5 {
// 			if len(result) > 0 {
// 				result = append(result, ' ')
// 			}
// 			result = append(result, block...)
// 			block = nil
// 			skipNext = true
// 		}
// 	}

// 	if len(block) > 0 {
// 		if len(result) > 0 {
// 			result = append(result, ' ')
// 		}
// 		result = append(result, block...)
// 	}

// 	result = append(result, '\n')
// 	return string(result)
// }
