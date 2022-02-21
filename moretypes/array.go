package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println("a:", a)
	fmt.Println(a[0])
	fmt.Println(a[1])
	// fmt.Println(a[2]) // out of bounds error

	b := [2]string{"hello", "world"}
	fmt.Println("b:", b)
}
