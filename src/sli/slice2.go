package main

import "fmt"

func main() {
	str := []string{"hello", "", "", "world"}
	res := nonempty(str)
	fmt.Println(res)
}

func nonempty(str []string) []string {
	i := 0
	for _, c := range str {
		if c != "" {
			str[i] = c
			i++
		}
	}
	return str[:i]
}
