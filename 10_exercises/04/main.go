package main

import "fmt"

var small int32
var big int32

func main() {
	fmt.Print("Please enter a small number here: ")
	fmt.Scan(&small)
	fmt.Print("Please enter a large number now: ")
	fmt.Scan(&big)
	fmt.Println("This is the remainder of the large number divided by the small number:", big%small)
}
