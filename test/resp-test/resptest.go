package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input := "$7\r\nRajveer\r\n"
	reader := bufio.NewReader(strings.NewReader(input))

	b, _ := reader.ReadByte()
	if b != '$' {
		fmt.Println("Invalid data type")
		os.Exit(1)
	}

	size, _ := reader.ReadByte()
	fmt.Println(size)
	strSize, _ := strconv.ParseInt(string(size), 10, 64)

	reader.ReadByte()
	reader.ReadByte() //move the reader ahead by one to consume /r and /n respectively

	name := make([]byte, strSize)
	reader.Read(name)

	fmt.Println(string(name))
}
