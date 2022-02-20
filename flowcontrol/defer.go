package main

import "fmt"

func main() {
	// main()のreturnのタイミングまで実行を遅延
	defer fmt.Print("太郎 ")
	defer fmt.Print("二郎 ")
	defer fmt.Print("三郎 ")
	fmt.Print("hello ")
}
