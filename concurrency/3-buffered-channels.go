package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// fatal error: all goroutines are asleep - deadlock!
	// ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
