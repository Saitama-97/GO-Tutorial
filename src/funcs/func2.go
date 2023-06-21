/**
 * @Time    : 2023/6/14 14:45
 * @File    : func2.go
 * @Project : tutorial
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc	: 匿名函数作为回调函数
 */

package main

import "fmt"

func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

func main() {
	func(data int) {
		fmt.Println("hello", data)
	}(100)

	f := func(data int) {
		fmt.Println("hello", data)
	}
	f(200)

	visit([]int{1, 2, 3}, func(v int) {
		fmt.Println(v)
	})
}
