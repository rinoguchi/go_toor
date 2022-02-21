package main

import "fmt"

func main() {
	brothers := []string{"taro", "jiro", "saburo"}
	for i, v := range brothers {
		fmt.Println(i, v)
	}
	for _, v := range brothers {
		fmt.Println(v)
	}
}
