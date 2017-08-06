package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func inputName() string {
	fmt.Println("Enter the name of the file:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return scanner.Text()
}
func lineCounter(fileName string) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := make([]byte, stat.Size())
	_, err = file.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	lineSep := []byte{'\n'}
	fmt.Println(bytes.Count(buf, lineSep), fileName)
}
func main() {
	fileName := inputName()
	lineCounter(fileName)
}
