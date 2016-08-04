package main

import "fmt"

func main() {
	half := func(z int) (int, bool) {
		return z / 2, z%2 == 0
	}
	h, even := half(5)
	fmt.Println(h, even)
}
