// +build no-build OMIT

package main

import "github.com/akwick/go-tutorial/pic"

type Image struct{}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
