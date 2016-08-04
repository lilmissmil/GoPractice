package main

import "fmt"

func main() {
	h, even := half(5)
	fmt.Println(h, even)
	fmt.Println(half(1)) // 0 false
	fmt.Println(half(2)) // 1 true
}

func half(z int) (int, bool) {
	return z / 2, z%2 == 0
}
