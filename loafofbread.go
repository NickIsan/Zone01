package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	var result []byte
	var block []byte
	skipNext := false

	i := 0
	for i < len(str) {
		ch := str[i]
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
