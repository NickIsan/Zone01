package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for first := 99; first >= 0; first-- {
		for second := 98; second >= 0; second-- {

			tens := first / 10
			ones := first % 10

			z01.PrintRune('0' + rune(tens))
			z01.PrintRune('0' + rune(ones))

			z01.PrintRune(' ')

			tens = second / 10
			ones = second % 10

			z01.PrintRune('0' + rune(tens))
			z01.PrintRune('0' + rune(ones))

			if first != 1 || second != 0 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
