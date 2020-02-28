package main

import (
	"fmt"
	"io"
	"os"
)

type ConsoleWriter struct{}

func main() {
	fileName := os.Args[1]

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err.Error())
	}

	var cw ConsoleWriter

	io.Copy(cw, file)

}

func (cw ConsoleWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}
