package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Printf("(%v, %T)\n", s, s)

	s, ok := i.(string)
	fmt.Printf("(%v, %T)\n", s, s)
	fmt.Printf("(%v, %T)\n", ok, ok)

	f, ok := i.(float64)
	fmt.Printf("(%v, %T)\n", f, f)
	fmt.Printf("(%v, %T)\n", ok, ok)

	// f := i.(float64)
	// fmt.Printf("(%v, %T)\n", f, f)
}
