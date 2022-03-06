package main

import "fmt"

type Greeter interface {
	Greet()
	Type() string
}

type Human struct {
	Name string
}

type Monkey struct {
	Name string
}

func (human Human) Greet() {
	fmt.Printf("Hello, my name is %v.\n", human.Name)
}

func (human Human) Type() string {
	return "Human"
}

func (monkey Monkey) Greet() {
	fmt.Printf("Wookey, wookakey wookkkey ukkety %v.\n", monkey.Name)
}

func (monkey Monkey) Type() string {
	return "Monkey"
}

func main() {
	var greeter1 Greeter = Human{"Taro"}
	println(greeter1.Type())
	greeter1.Greet()

	var greeter2 Greeter = Monkey{"Jiro"}
	println(greeter2.Type())
	greeter2.Greet()

}
