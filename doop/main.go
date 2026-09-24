package main

import (
	"os"
)

func atoi(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}

	sign := int64(1)
	start := 0

	if s[0] == '-' {
		sign = -1
		start = 1
	}

	if start == len(s) {
		return 0, false
	}

	var n int64

	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}

		n = n*10 + int64(s[i]-'0')
	}

	return n * sign, true
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	a, ok := atoi(os.Args[1])
	if !ok {
		return
	}

	b, ok := atoi(os.Args[3])
	if !ok {
		return
	}

	switch os.Args[2] {
	case "+":
		print(a + b)

	case "-":
		print(a - b)

	case "*":
		print(a * b)

	case "/":
		if b == 0 {
			print("No division by 0\n")
			return
		}
		print(a / b)

	case "%":
		if b == 0 {
			print("No modulo by 0\n")
			return
		}
		print(a % b)

	default:
		return
	}
}
