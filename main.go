package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := os.Args[1]

	fmt.Printf("%b\n", os.O_RDWR|os.O_CREATE)

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	mw := io.MultiWriter(os.Stdout, file)
	if isInputFromPipe() {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(os.Stdin)
		if err != nil {
			panic(err)
		}
		mw.Write((buf.Bytes()))
	} else {
		r := bufio.NewReader(os.Stdin)
		input, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}
		mw.Write([]byte(input))
	}

}

func isInputFromPipe() bool {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	return fileInfo.Mode()&os.ModeCharDevice == 0
}
