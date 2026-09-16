package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	name := filepath.Base(os.Arg[0])
	fmt.Println(name)
}
