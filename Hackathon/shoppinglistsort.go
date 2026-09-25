package piscine

func ShoppingListSort(slice []string) []string {
	for i := 0; i < len(slice); i++ {
		min := i

		for j := i + 1; j < len(slice); j++ {
			if len(slice[j]) < len(slice[min]) {
				min = j
			}
		}

		slice[i], slice[min] = slice[min], slice[i]
	}

	return slice
}
