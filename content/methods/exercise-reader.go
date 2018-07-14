// +build no-build OMIT

package main

import "github.com/akwick/go-tutorial/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
