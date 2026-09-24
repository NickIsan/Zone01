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

func printString(s string) {
	os.Stdout.Write([]byte(s))
}

func printNumber(n int64) {
	if n == 0 {
		printString("0")
		return
	}

	if n < 0 {
		printString("-")
		n = -n
	}

	var digits [20]byte
	i := 0

	for n > 0 {
		digits[i] = byte(n%10) + '0'
		n = n / 10
		i++
	}

	for i > 0 {
		i--
		os.Stdout.Write([]byte{digits[i]})
	}
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
		printNumber(a + b)

	case "-":
		printNumber(a - b)

	case "*":
		printNumber(a * b)

	case "/":
		if b == 0 {
			printString("No division by 0\n")
			return
		}

		printNumber(a / b)

	case "%":
		if b == 0 {
			printString("No modulo by 0\n")
			return
		}

		printNumber(a % b)

	default:
		return
	}
}
