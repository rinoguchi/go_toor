package main

import (
	"fmt"
	"strconv"
)

func add(x, y int, z string) int {
	i, _ := strconv.Atoi(z)
	return x + y + i
}

func main() {
	fmt.Println(add(12, 34, "56"))
}
