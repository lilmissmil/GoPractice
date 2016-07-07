package main

import "fmt"

// %d is decimal, the - is a dash just to separate it from %b, which is binary
// The \n is an escape character that will create a new line
// The # will add 0x in front of x to produce another way to represent hexadecimal
// The \t will create a space or a tab between the variables/strings
// %q prints the actual value from UTF-8

func main() {
	//fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	//fmt.Printf("%d - %b - %#x \n", 42, 42, 42)
	//fmt.Printf("%d - %b - %#X \n", 42, 42, 42)
	//fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
