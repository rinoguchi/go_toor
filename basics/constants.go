package main

import "fmt"

const pi = 3.14

func main() {
	const world = "World"
	fmt.Println("Hello", world)

	radius := 4.00
	fmt.Printf("Area of a circle with radius %v cm is %v\n", radius, radius*radius*pi)
}
