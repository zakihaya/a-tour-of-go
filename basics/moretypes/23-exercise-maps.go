package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	m := make(map[string]int)

	for _, v := range words {
		m[v] += 1
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
