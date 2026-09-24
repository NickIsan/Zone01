package piscine

func StringToIntSlice(str string) []int {
	var numbers []int

	for _, letter := range str {
		numbers = append(numbers, int(letter))
		// numbers[i] = int(letter)
	}
	return numbers
}
