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

/*  for pass := 0; pass < 4; pass++ {
        for i := 0; i < 4; i++ {
            if n[i] > n[i+1] {
                n[i], n[i+1] = n[i+1], n[i]
            }
        }
    }




	for swapped {
    swapped = false

    for i := 0; i < len(n)-1; i++ {
        if n[i] > n[i+1] {
            n[i], n[i+1] = n[i+1], n[i]
            swapped = true
        }
    }
}
*/
