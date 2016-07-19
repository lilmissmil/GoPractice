package main

import "fmt"

func main() {
	if !true {
		fmt.Println("This is false.")
	}
	if !false {
		fmt.Println("This is true.")
	}

	b := true
	if food := "Chocolate"; b {
		fmt.Println(food)
	}
	if false {
		fmt.Println("First statement")
	} else if false {
		fmt.Println("Second statement")
	} else {
		fmt.Println("Third statement")
	}
}
