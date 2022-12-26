package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var m2 map[string]int

func main() {
	m = make(map[string]Vertex)
	m["Bells Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bells Labs"])

	m2 = make(map[string]int)
	m2["a"] = 1
	fmt.Println(m2)
}
