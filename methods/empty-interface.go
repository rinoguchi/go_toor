package main

import "fmt"

func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("this is int")
	case string:
		fmt.Println("this is string")
	default:
		fmt.Println("this is other")
	}

}

func main() {
	fmt.Println("=========== 基本 ===========")
	var unknown1 interface{}
	fmt.Printf("%v, %T\n", unknown1, unknown1)
	unknown1 = "ABC"
	fmt.Printf("%v, %T\n", unknown1, unknown1)
	unknown1 = 123
	fmt.Printf("%v, %T\n", unknown1, unknown1)

	fmt.Println("=========== 型チェック ===========")
	var unknown2 interface{}
	str, ok := unknown2.(string)
	fmt.Println(str, ok)
	unknown2 = "ABC"
	str, ok = unknown2.(string)
	fmt.Println(str, ok)
	unknown2 = 123
	str, ok = unknown2.(string)
	fmt.Println(str, ok)
	// str = unknown2.(string) // => `panic: interface conversion: interface {} is int, not string` occurrs

	fmt.Println("=========== 型スイッチ ===========")
	var unknown3 interface{}
	printType((unknown3))
	unknown3 = "ABC"
	printType((unknown3))
	unknown3 = 123
	printType((unknown3))
}
