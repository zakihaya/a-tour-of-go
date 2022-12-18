package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	vertex := Vertex{1, 2}
	fmt.Println(vertex)
	fmt.Println(vertex.X)
	fmt.Println(vertex.Y)
}
