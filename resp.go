package main

import (
	"bufio"
	"io"
)

const (
	STRING  = '+'
	ERROR   = '-'
	INTEGER = ':'
	BULK    = '$'
	ARRAY   = '*'
)

type value struct {
	typ   string
	str   string
	num   int
	bulk  string
	array []value
}

type Resp struct {
	reader *bufio.Reader
}

/*  Initialize a structure Resp that has "field"(basically an entry)
    called reader that is a pointer to a struct Reader declared in bufio*/

func NewResp(rd io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(rd)}
}

/*  Declare a new function called NewResp that takes in a value of datatype
    Reader defined in "io" and returns a pointer to a variable of
    datatype Resp, in the function the memory address of a newly
    initialized buffered reader is returned.*/

func (r *Resp) readLine() (line []byte, n int, err error) {
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}

		n++
		line = append(line, b)
		if len(line) >= 2 && line[len(line)-2] == '\r' {
			break
		}
	}
	return line[:len(line)-2], n, err //weird ass syntax
}

/*  Take a pointer to a var of datatype Resp and return a byte array,
the line is calculated by iteratively reading the byte array byte
by byte into another byte array and checking if the '\r' has
been read or not.
The reason it wouldnt break if a '\r' arrives in the payload is
because the bulk string which will be later defined will be read
using the length and not the delimiters.*/
