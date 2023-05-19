package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fName := os.Args[1]

	fp, err := os.Open(fName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, fp)
}
