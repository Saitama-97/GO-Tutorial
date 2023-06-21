/**
 * @Time    : 2023/6/14 15:56
 * @File    : func3.go
 * @Project : tutorial
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : 匿名函数实现操作封装
 */

package main

import (
	"flag"
	"fmt"
)

var skillParam = flag.String("skill", "", "skill to perform")
var map1 map[string]int = map[string]int{
	"one": 1,
	"two": 2,
}

func main() {
	flag.Parse()

	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("solier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
}
