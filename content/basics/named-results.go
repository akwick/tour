// +build OMIT

package main

import "fmt"

func add42(x, y int) (a, b int) {
	a = x + 42
	b = y + 42
	return
}

func main() {
	fmt.Println(add42(0, 1))
}