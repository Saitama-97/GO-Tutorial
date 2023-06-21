/**
 * @Time    : 2023/6/16 14:55
 * @File    : func7.go
 * @Project : tutorial
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : 函数可被看作第一类值（first-class values）
 */

package main

import (
	"fmt"
	"strings"
)

func square(n int) int {
	return n * n
}

func negative(n int) int {
	return -n
}

func product(m, n int) int {
	return m * n
}

func add1(r rune) rune {
	return r + 1
}

func main() {
	f := square
	a, b := 2, 3
	fmt.Println(f(a), square(a), negative(b), product(a, b))
	// map function ，returns a copy of the string s with all its characters modified according to the mapping function
	// ascii code add 1
	fmt.Println(strings.Map(add1, "HAL-9000"))
	fmt.Println(strings.Map(add1, "VMS"))
	fmt.Println(strings.Map(add1, "admin"))
}
