package main

import (
	"fmt"
)

type MyReader struct{}

// Add a Read([]byte) (int, error) method to MyReader.
func (reader MyReader) Read(buffer []byte) (int, error) {
	bufferLength := len(buffer)

	for i := 0; i < bufferLength; i++ {
		buffer[i] = 'A'
	}

	return bufferLength, nil
}

func main() {
	// reader.Validate(MyReader{})
	reader := MyReader{}
	b := make([]byte, 8)
	n, err := reader.Read(b)
	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
	fmt.Printf("b[:n] = %q\n", b[:n])
}
