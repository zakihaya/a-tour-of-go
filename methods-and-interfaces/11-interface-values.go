package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	fmt.Printf("pointer main i: %v\n", &i)
	describe(i)
	describe2(&i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
	fmt.Printf("pointer i: %v\n", &i)
}

func describe2(i *I) {
	fmt.Printf("pointer *i: %v\n", i)
	(*i).M()
}
