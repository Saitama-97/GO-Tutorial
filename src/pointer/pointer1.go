package main

import "fmt"

func main() {
	x := 1
	adr := &x
	fmt.Println(x)
	fmt.Println(adr)
	fmt.Println(*adr)
}
