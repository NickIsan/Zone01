package piscine

func ReverseMenuIndex(menu []string) []string {
	result := make([]string, 4)
	for i := 0; i < 4; i++ {
		result[i] = menu[i]
	}
	left := 0
	right := len(menu) - 1

	for left < right {

		menu[left], menu[right] = menu[right], menu[left]

		left++
		right--
	}

	return menu
}
