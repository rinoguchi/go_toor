package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %v\n", i, f, b, s)
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
