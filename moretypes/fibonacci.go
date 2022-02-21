package main

import "fmt"

func fibonacci() func() int {
	x := 0
	y := 0
	return func() int {
		if y == 0 {
			x, y = 0, 1
		} else {
			x, y = y, x+y
		}
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
