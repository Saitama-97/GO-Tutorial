package main

import "fmt"

/**
 * @Time    : 2023/7/11 10:29
 * @File    : method1.go
 * @Project : tutorial
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : GO 语言中的方法：包含了接受者的函数
 */

type Square struct {
	length float32
}

func main() {
	s1 := Square{length: 3.0}
	var area float32
	area = s1.getArea()
	fmt.Printf("Length-%.2f, Area-%.2f\n", s1.length, area)
}

// getArea 一个方法，指定了接受者为Square类型的一个对象
func (s Square) getArea() float32 {
	return s.length * s.length
}
