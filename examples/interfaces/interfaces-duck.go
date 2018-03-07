package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

type customReader struct {
	content  string
	position int
}

func (r *customReader) Read(buf []byte) (int, error) {
	if r.position >= len(r.content) {
		return 0, io.EOF
	}

	n := copy(buf, r.content[r.position:])
	r.position += n
	return n, nil
}

// START OMIT
func main() {
	// struct with a Read method
	v := &customReader{
		content: "Some text",
	}

	// ioutil.ReadAll(r io.Reader) ([]byte, error)
	bytes, _ := ioutil.ReadAll(v) // HL
	fmt.Printf("bytes = %#v", bytes)
}

// END OMIT
