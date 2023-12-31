package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Failed to open the file:%v\n", err.Error())
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
