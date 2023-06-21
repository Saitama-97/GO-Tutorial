package main

import "fmt"

func typedTwoValues() (int, int) {
	return 1, 2
}

func namedRetValues() (a, b int) {
	a = 1
	b = 2
	return
}

func fire() {
	fmt.Println("fire in the hole")
}

func main() {
	a, b := typedTwoValues()
	fmt.Printf("%d %d\n", a, b)

	c, d := namedRetValues()
	fmt.Printf("%d %d\n", c, d)

	var new_fire func()
	new_fire = fire
	new_fire()
}
