package piscine

func StringToIntSlice(str string) []int {
	numbers := make([]int, len(str))
	for _, letter := range str {
		numbers = append(numbers, int(letter))
		// numbers[i] = int(letter)
	}
	return numbers
}
