package main

import "fmt"

func main() {
	// makeを使って初期化
	var m1 map[string]int
	m1 = make(map[string]int)
	m1["taro"] = 170
	m1["jiro"] = 180
	fmt.Println("m1:", m1)

	// makeを使わず初期化
	m2 := map[string]int{
		"saburo": 190,
		"shiro":  200,
	}
	fmt.Println("m2:", m2)

	// 要素を取得
	fmt.Println("m2[\"saburo\"]:", m2["saburo"])

	// 要素の削除
	delete(m2, "shiro")
	fmt.Println("m2:", m2)

	// 要素の追加
	m2["goro"] = 210
	fmt.Println("m2:", m2)

	// 要素の存在確認
	elem, ok := m2["goro"]
	fmt.Println(elem, ok)
	elem, ok = m2["rokuro"]
	fmt.Println(elem, ok)
}
