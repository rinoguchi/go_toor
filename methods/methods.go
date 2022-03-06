package main

import (
	"fmt"
)

// 長方形
type Rectangle struct {
	X, Y float64
}

// 面積計算 値レシーバ
func (r Rectangle) Area() float64 {
	return r.X * r.Y
}

// 面積計算 ポインタレシーバ
func (r *Rectangle) AreaR() float64 {
	return r.X * r.Y
}

// 拡大縮小 値レシーバ
func (r Rectangle) Scale(f float64) {
	r.X = r.X * f
	r.Y = r.Y * f
}

// 拡大縮小 ポインタレシーバ
func (r *Rectangle) ScaleR(f float64) {
	r.X = r.X * f
	r.Y = r.Y * f
}

func main() {
	fmt.Println("===== 値レシーバ × 値 =====")
	r1 := Rectangle{3.0, 5.0}
	fmt.Printf("area: %v\n", r1.Area())
	r1.Scale(2)
	fmt.Printf("area: %v\n", r1.Area())

	fmt.Println("===== ポインタレシーバ × 値 =====")
	r2 := Rectangle{3.0, 5.0}
	fmt.Printf("area: %v\n", r2.AreaR())
	r2.ScaleR(2)
	fmt.Printf("area: %v\n", r2.AreaR())

	fmt.Println("===== 値レシーバ × ポインタ =====")
	p1 := &Rectangle{3.0, 5.0}
	fmt.Printf("area: %v\n", p1.Area())
	p1.Scale(2)
	fmt.Printf("area: %v\n", p1.Area())

	fmt.Println("===== ポインタレシーバ × ポインタ =====")
	p2 := &Rectangle{3.0, 5.0}
	fmt.Printf("area: %v\n", p2.AreaR())
	p2.ScaleR(2)
	fmt.Printf("area: %v\n", p2.AreaR())
}
