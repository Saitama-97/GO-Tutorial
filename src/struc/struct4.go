package main

import "fmt"

type Wheel struct {
	Size int
}

type Car struct {
	Wheel
	Engine struct {
		Power int
		Type  string
	}
}

func main() {
	//c1 := Car{
	//	Wheel:  Wheel{225},
	//	Engine: Engine{200, "1.4T"},
	//}

	c1 := Car{
		Wheel: Wheel{225},
		Engine: struct {
			Power int
			Type  string
		}{
			Power: 200,
			Type:  "1.4T",
		},
	}

	fmt.Printf("%v\n", c1)
	fmt.Printf("%+v\n", c1)
}
