package main

import "golang.org/x/tour/tree"
import "fmt"

// type Tree struct {
//     Left  *Tree
//     Value int
//     Right *Tree
// }

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return true
}

func main() {
	ch := make(chan int)
	t := tree.New(1)
	
	fmt.Println(ch<-)
	
	Walk(t, ch)
	
	fmt.Println("===")
	fmt.Println(t)
	fmt.Println(t.Left)
	fmt.Println(t.Right)
	fmt.Println(t.Value)
	fmt.Println(t.Left)
	fmt.Println(t.Left.Left)
	fmt.Println(t.Left.Left.Left)
	fmt.Println(t.Left.Left.Left.Left == nil)
	fmt.Println(t.Left.Left.Left.Right)
	fmt.Println(t.Left.Left.Left.Value)
}
