package main

import "fmt"

type Cat struct {
	Color string
	Name  string
}

type BlackCat struct {
	Cat
}

func NewCatByName(name string) *Cat {
	cat := Cat{
		Name: name,
	}
	return &cat
}

func NewBlackCat(color string) *BlackCat {
	cat := BlackCat{}
	cat.Color = color
	return &cat
}

func struct2() {
	//addr := NewCatByName("Jerry")
	addr := NewBlackCat("Black")
	fmt.Println(addr.Color)
}
