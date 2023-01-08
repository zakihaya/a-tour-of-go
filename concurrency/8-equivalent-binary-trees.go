package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)
	values1 := [10]int{}
	for i := 0; i < 10; i++ {
		n := <-ch1
		values1[i] = n
	}

	ch2 := make(chan int)
	go Walk(t2, ch2)
	values2 := [10]int{}
	for i := 0; i < 10; i++ {
		n := <-ch2
		values2[i] = n
	}

	return values1 == values2
}

func main() {
	ch := make(chan int)
	t := tree.New(1)

	go Walk(t, ch)

	for i := 0; i < 10; i++ {
		n := <-ch
		fmt.Println(n)
	}

	fmt.Println("===")
	fmt.Println("Same(tree.New(1), tree.New(1))")
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println("===")
	fmt.Println("Same(tree.New(1), tree.New(2))")
	fmt.Println(Same(tree.New(1), tree.New(2)))
	// fmt.Println("===")
	// fmt.Println(t)
	// fmt.Println(t.Left)
	// fmt.Println(t.Right)
	// fmt.Println(t.Value)
	// fmt.Println(t.Left)
	// fmt.Println(t.Left.Value)
	// fmt.Println(t.Left.Left)
	// fmt.Println(t.Left.Left.Left)
	// fmt.Println(t.Left.Left.Left.Left == nil)
	// fmt.Println(t.Left.Left.Left.Right)
	// fmt.Println(t.Left.Left.Left.Value)
}
