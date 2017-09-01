package main

import (
	"fmt"
	"os"
	"io"
)

const bufSize 	int 	= 10
const filePath 	string 	= "./resources/readme.txt"

func print(channel chan string) {
	var (
		isopen 	bool 	= true
		chr		string 	= ""
	)
	for isopen {
		chr, isopen = <- channel
		fmt.Print(chr)
	}
}

func readFile(channel chan string) {
	file, fileOpenError := os.Open(filePath)
	defer file.Close()

	err := fileOpenError
	buf := make([]byte, bufSize)
	for err != io.EOF {
		bytesRead, fileReadError := file.Read(buf)
		err = fileReadError
		channel <- string(buf[:bytesRead])
	}
	close(channel)
}

func main() {

	strChannel := make(chan string)
	go readFile(strChannel)
	print(strChannel)
}
