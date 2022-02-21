package main

import "fmt"

func main() {
	// sliceの初期化
	brothers := []string{"taro", "jiro", "saburo", "siro", "goro"}
	fmt.Println("brothers:", brothers)
	fmt.Println("len=", len(brothers), "cap=", cap(brothers))

	// appendした時のlenとcap確認
	brothers = append(brothers, "rokuro", "shichiro")
	fmt.Println("brothers:", brothers)
	fmt.Println("len=", len(brothers), "cap=", cap(brothers))

	// サブsliceを取り出した時のlenとcap確認
	brothers = brothers[0:3]
	fmt.Println("brothers:", brothers)
	fmt.Println("len=", len(brothers), "cap=", cap(brothers))

	// zero slice
	var zeroSlice []int
	fmt.Println("zeroSlice", zeroSlice)
	fmt.Println("len=", len(zeroSlice), "cap=", cap(zeroSlice))
	fmt.Println("zeroSlice == nil:", zeroSlice == nil)

	// makeでlenとcapを指定して初期化
	sisters := make([]string, 2, 5)
	fmt.Println("sisters:", sisters)
	fmt.Println("len=", len(sisters), "cap=", cap(sisters))
	sisters[0] = "ichiko"
	sisters[1] = "niko"
	// sisters[2] = "sanko" index out of range error
	fmt.Println("sisters:", sisters)
	fmt.Println("len=", len(sisters), "cap=", cap(sisters))
}
