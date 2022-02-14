package main

import "fmt"

func divide(num, denominator int) (quotient, remainder int) {
	quotient = num / denominator
	remainder = num % denominator
	return
}

func main() {
	fmt.Println(divide(35, 8))
}
