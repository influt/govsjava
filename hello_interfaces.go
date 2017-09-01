package main

import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
)

type reader interface {
	readFile(chan string)
}

type bufFileReader 	struct {}
type utilFileReader struct {}

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

func (bufReader bufFileReader) readFile(channel chan string) {
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

func (fileReader utilFileReader) readFile(channel chan string) {
	content, err := ioutil.ReadFile(filePath)
	if err == nil {
		channel <- string(content)
	}
	close(channel)
}

func main() {
	var fileReader reader = utilFileReader{}
	
	strChannel := make(chan string)
	
	go fileReader.readFile(strChannel)
	print(strChannel)
}
