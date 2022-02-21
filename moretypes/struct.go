package main

import "fmt"

type Physique struct {
	Height int
	Weight int
}

type Human struct {
	Name     string
	Physique Physique
}

func main() {
	fmt.Println("==========taro==========")
	taroPhysique := Physique{Height: 170, Weight: 60}
	taro := Human{Name: "taro", Physique: taroPhysique}
	fmt.Println("taro:", taro)
	fmt.Println("taro.Name:", taro.Name)
	fmt.Println("taro.Physique:", taro.Physique)
	taroP := &taro
	fmt.Println("taroP:", taroP)
	fmt.Println("taroP.Name:", taroP.Name)
	fmt.Println("*taroP.Name:", taroP.Name)

	fmt.Println("==========jiro==========")
	jiro := Human{Name: "jiro", Physique: Physique{Height: 180, Weight: 90}}
	fmt.Println("jiro:", jiro)

	fmt.Println("==========saburo==========")
	saburo := Human{Name: "saburo"}
	fmt.Println("saburo:", saburo)
	saburo.Physique.Height = 200
	saburo.Physique.Weight = 100
	fmt.Println("saburo:", saburo)
}
