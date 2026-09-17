package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	// numbers := make([]int, 0)
	numbers := []int{}
	// var numbers []int
	for i := min; i < max; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}
