package main

import (
	"fmt"
	"reflect"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var arr [3]int
	for i, v := range arr {
		fmt.Printf("%d-%d\n", i, v)
	}
	symbol := [...]string{USD: "U", EUR: "E", GBP: "G", RMB: "R"}
	fmt.Println(reflect.TypeOf(symbol[3]), symbol[3])
}
