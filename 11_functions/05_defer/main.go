package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world()
	hello()

}

// defer basically leaves that you deferred to the end right before the main function closes
// useful for opening a file and making sure it gets closed to save the computer's memory etc.
