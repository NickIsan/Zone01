package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args {
		fmt.Print("os.Args[", index, "] == ", arg)
		fmt.Println()
	}
	fmt.Println()
}
