package main

import "fmt"

func main() {
	i := 100

	var p *int = &i // p := &i と同等

	fmt.Println(p)
	fmt.Println(*p)

	*p = 200
	fmt.Println(p)
	fmt.Println(*p)

	*p = *p + 100
	fmt.Println(p)
	fmt.Println(*p)
}
