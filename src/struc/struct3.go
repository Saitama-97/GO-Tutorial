package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b      int
	c      float32
	int    // anonymous
	innerS // anonymous
}

func struct3() {
	outer := new(outerS)
	//fmt.Println(outer)
	outer.b = 1
	outer.c = 2
	outer.int = 3
	outer.in1 = 4
	outer.in2 = 5
	fmt.Println(outer)

	outer2 := outerS{
		b:      11,
		c:      22,
		int:    33,
		innerS: innerS{44, 55},
	}

	fmt.Println(outer2)
}
