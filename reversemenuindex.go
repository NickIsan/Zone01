package piscine

func ReverseMenuIndex(menu []string) []string {
	left := 0
	right := len(menu) - 1

	for left < right {

		menu[left], menu[right] = menu[right], menu[left]

		left++
		right--
	}

	return menu
}
