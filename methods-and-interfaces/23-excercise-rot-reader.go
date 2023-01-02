package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(bytes []byte) (int, error) {
	buffer := make([]byte, len(bytes))
	bufferLen, err := reader.r.Read(buffer)
	if err != nil {
		return 0, err
	}

	for i := 0; i < bufferLen; i++ {
		if 'A' <= buffer[i] && buffer[i] <= 'M' {
			bytes[i] = buffer[i] + 13
		} else if 'a' <= buffer[i] && buffer[i] <= 'm' {
			bytes[i] = buffer[i] + 13
		} else {
			bytes[i] = buffer[i] - 13
		}
	}
	return bufferLen, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
