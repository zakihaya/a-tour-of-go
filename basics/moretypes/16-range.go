package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 128}

func main() {
	// range is implementation of iterator
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
