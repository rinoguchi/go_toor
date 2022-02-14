package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 4, 7
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	var s string = fmt.Sprint(z)
	// var s string = string(z) => not working expectedly
	fmt.Println(x, y, f, z, s)
}
