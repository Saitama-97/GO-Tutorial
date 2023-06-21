package main

import (
	"fmt"
	"reflect"
)

func slice1() {
	runes := make([]rune, 0)
	for _, r := range "Hello, World!" {
		runes = append(runes, r)
	}
	fmt.Println(len(runes), runes)
	for i, r := range runes {
		fmt.Println(i, string(r))
	}
	r1 := 'a'
	var r2 rune = 'a'
	fmt.Println(r1)
	fmt.Println(reflect.TypeOf(r1))
	fmt.Println(reflect.TypeOf(r2))
}
