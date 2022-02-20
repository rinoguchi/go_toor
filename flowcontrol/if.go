package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sqrt(10.0))
	fmt.Println(sqrt(-10.0))
	fmt.Println(pow(2, 4))
	fmt.Println(pow(2, 8))
	fmt.Println(pow(2, 12))
}

func sqrt(v float64) string {
	// simple if
	if v < 0 {
		return sqrt(-v) + "i"
	}
	return fmt.Sprint(math.Sqrt(v))
}

func pow(x, n float64) string {
	// if with initialize
	if v := math.Pow(x, n); v < 100 {
		return fmt.Sprintf("%v: less than 100", v)
	} else if v < 1000 {
		return fmt.Sprintf("%v: less than 1000", v)
	} else {
		return fmt.Sprintf("%v: over 1000", v)
	}
}
