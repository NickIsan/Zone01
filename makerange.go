package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	size := max - min

	numbers := make([]int, size)

	for i := 0; i < size; i++ {
		for j := min; j < max; j++ {
			if j-min == i {
				numbers[i] = j
			}
		}
	}
	return numbers
}
