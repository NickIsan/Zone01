package main

import (
	"io"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for i := 1; i < len(os.Args); i++ {
		file, err := os.Open(os.Args[i])
		if err != nil {
			os.Stderr.Write([]byte("ERROR: " + err.Error() + "\n"))

			if len(os.Args) == 2 {
				os.Exit(1)
			}

			return
		}

		io.Copy(os.Stdout, file)

		file.Close()
	}
}
