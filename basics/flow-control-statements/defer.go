package main

import "fmt"

func main() {
	defer fmt.Println("World")

	n := 3
	defer fmt.Println(n)

	fmt.Println("Hello")

	n = 4

	// result
	// ======
	// Hello
	// 3
	// World
}
