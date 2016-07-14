package main

import "fmt"

//The operator, &, placed before the operand will give us the memory address

func main() {

	a := 42

	fmt.Println("a -", a)
	fmt.Println("memory address -", &a)
	fmt.Printf("decimal - %d \n", &a)

}
