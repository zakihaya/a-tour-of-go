package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	currentIndex := 0
	n1 := 0
	n2 := 0
	return func() int {
		currentIndex += 1
		if currentIndex == 1 {
			return 0
		}
		if currentIndex == 2 {
			n2 = 1
			return 1
		}
		n := n1 + n2
		n1 = n2
		n2 = n

		return n
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
