package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2} // has type vertex
	v2 = Vertex{X: 1} // Y: 0 is implicit
	v3 = Vertex{}     // X: 0 and Y: 0
	p  = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
	fmt.Printf("%T %T %T %T", v1, p, v2, v3)
}
