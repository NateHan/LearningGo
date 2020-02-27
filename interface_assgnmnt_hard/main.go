package main

import (
	"fmt"
	"os"
	"io"
)

type consolePrinter struct {}

type ConsoleReader struct{}

type ConsoleWriter struct{
	File *os.File
}


func main() {
	fileName := os.Args[1]

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err.Error())
	}

	

	io.Copy( , file)

}

func (cr ConsoleReader) Read(p []byte) (n int, err error) {

	fmt.Println(string(p))
	return len(p), nil
}

func (cw ConsoleWriter) Write(p []byte) (n int, err error) {

}
