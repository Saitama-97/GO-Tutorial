package main

type A struct {
	a int
}

type B struct {
	b int
}

type C struct {
	A
	B
}

func struct5() {
	c1 := new(C)
	c1.a = 1
}
