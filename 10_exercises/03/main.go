package main

import "fmt"

var name string

func main() {
	fmt.Print("Please print your name here:")
	fmt.Scan(&name)
	fmt.Println("Hello,", name)

}
