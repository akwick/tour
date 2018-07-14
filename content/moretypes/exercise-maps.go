// +build OMIT

package main

import (
	"github.com/akwick/go-tutorial/wc"
)

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
