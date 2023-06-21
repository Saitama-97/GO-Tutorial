/**
 * @Time    : 2023/6/14 16:13
 * @File    : func4.go
 * @Project : tutorial
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : 函数类型实现接口
 */

package main

import (
	"fmt"
	"reflect"
)

type Invoker interface {
	Call(interface{})
}

type Struct struct {
}

func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

func main() {
	var a int
	fmt.Println(a)
	//fmt.Println(a == nil)
	fmt.Println(reflect.DeepEqual(a, nil))
}
