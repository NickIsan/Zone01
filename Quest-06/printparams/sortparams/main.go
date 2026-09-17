package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		for j := i + 1; j < len(os.Args); j++ {
			if os.Args[i] > os.Args[j] {
				os.Args[i], os.Args[j] = os.Args[j], os.Args[i]
			}
		}
	}

	for i := 1; i < len(os.Args); i++ {
		for _, r := range os.Args[i] {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
