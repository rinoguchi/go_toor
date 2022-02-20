package main

import "fmt"

func main() {
	simple()
	withoutPrePostProcess()
	infinitLoop()
}

// シンプルなfor文
func simple() {
	sum := 1
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("simple", sum)
}

// 初期化、後処理なし
func withoutPrePostProcess() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println("without pre post process", sum)
}

// 初期化、後処理なし
func infinitLoop() {
	sum := 1
	for {
		sum += sum

		if sum > 100000000 {
			break
		}
	}
	fmt.Println("loop", sum)
}
