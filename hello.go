package main

import (
	"fmt"
	"os"
	"io"
)

func main() {
	file, fileOpenError := os.Open("./resources/readme.txt")
	err := fileOpenError
	bufSize := 1024
	for i:=bufSize; err != io.EOF; i+=bufSize {
		buf := make([]byte, bufSize)
		_, fileReadError := file.Read(buf)
		err = fileReadError
		fmt.Printf("%s\n", string(buf))
	}
}
