package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			z01.PrintRune(rune(i/10) + '0')
			z01.PrintRune(rune(i%10) + '0')
			z01.PrintRune(' ')
			z01.PrintRune(rune(j/10) + '0')
			z01.PrintRune(rune(j%10) + '0')

			if i != 98 || j != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

/*
package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			for c := '0'; c <= '9'; c++ {
				for d := '0'; d <= '9'; d++ {

					first := (a-'0')*10 + (b-'0')
					second := (c-'0')*10 + (d-'0')

					if first < second {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(' ')
						z01.PrintRune(c)
						z01.PrintRune(d)

						if first != 98 || second != 99 {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}

	z01.PrintRune('\n')
}

*/
