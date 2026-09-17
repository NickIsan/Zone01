package piscine

func ConcatParams(args []string) string {
	result := ""

	for i := 0; i < len(args); i++ {
		result += args[i]

		if i < len(args)-1 {
			result += "\n"
		}
	}

	return result
}

/*
func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	size := len(args) -1

	for _, arg := ranges args {
		size += len(arg)

	result := make([byte], size)

	position := 0

	for i, arg := range args {
		for j := 0; j < len(arg); j++ {
			result[position] = arg[j]
			position++
		}

		if i < len(args)-1 {
			result[position] = '\n'
			position++
		}
	}
	}
	return string(result)
}
*/
