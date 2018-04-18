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

func main() {
	// START OMIT
	// custom struct with a Read method
	v := &customReader{
		content: "Some text",
	}

	// ioutil.ReadAll is from the stdlib and uses io.Reader
	bytes, _ := ioutil.ReadAll(v) // can use custom struct // HL
	// END OMIT
	fmt.Printf("bytes = %#v", bytes)
}
