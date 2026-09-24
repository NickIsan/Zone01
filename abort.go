package piscine

func Abort(a, b, c, d, e int) int {
	n := []int{a, b, c, d, e}

	for i := 0; i < 3; i++ {
		for j := i + 1; j < 5; j++ {
			if n[i] > n[j] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}

	return n[2]
}
