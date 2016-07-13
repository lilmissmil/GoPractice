package main

import "fmt"

// Constants can be in package scope or block scope depending on usage
// A constant is a simple unchanging value
// You can name more than one constant as shown below with Pi and Language

const (
	Pi       = 3.14
	Language = "Go"
)

// The above is also known as multiple initialization

func main() {
	const p string = "death & taxes"
	const q = 42
	fmt.Println("p -", p)
	fmt.Println("q -", q)
	fmt.Println("Pi -", Pi)
	fmt.Println("Language -", Language)

}
