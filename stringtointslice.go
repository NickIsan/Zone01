package piscine

func StringToIntSlice(str string) []int {
	numbers := make([]int, len(str))
	for i, letter := range str {
		numbers[i] = int(letter)
	}
	return numbers
}
