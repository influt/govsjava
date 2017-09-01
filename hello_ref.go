package main

import (
	"fmt"
	"os"
	"io"
)

var bufSize 	int 	= 1024
var filePath 	string 	= "./resources/readme.txt"

func main() {
	file, fileOpenError := os.Open(filePath)
	defer file.Close()

	err := fileOpenError
	buf := make([]byte, bufSize)
	for err != io.EOF {
		fmt.Printf("%s\n", string(buf))
		bytesRead, fileReadError := file.Read(buf)
		err = fileReadError
	}
}
