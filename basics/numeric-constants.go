package main

import "fmt"

const (
	big   = 1 << 100
	small = big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(small)
	// fmt.Println(big) => error: cannot use big (untyped int constant 1267650600228229401496703205376) as int value in argument to fmt.Println (overflows)
	fmt.Println(needInt(small))
	// fmt.Println(needInt(big)) => error: cannot use big (untyped int constant 1267650600228229401496703205376) as int value in argument to fmt.Println (overflows)
	fmt.Println(needFloat(small))
	fmt.Println(needFloat(big))
}
