package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i < 100; i++ {
		newZ := z - (z*z-x)/(2*z)
		fmt.Println(i, ":", newZ)
		if math.Abs(newZ-z) < 0.00000000001 {
			break
		}
		z = newZ
	}
	return z
}

func main() {
	fmt.Println("result:", Sqrt(2))
	fmt.Println("result:", Sqrt(3))
	fmt.Println("result:", Sqrt(4))
}
