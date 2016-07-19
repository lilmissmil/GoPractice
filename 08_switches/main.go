package main

import "fmt"

func main() {
	switch "John" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi", "John":
		fmt.Println("Wassup Medhi, or, err John")
		fallthrough
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have you no friends?")
	}

	myFriendsName := "Mar"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Wassup my friend with name of length 2")
	case myFriendsName == "Tim":
		fmt.Println("Wassup Tim")
	case myFriendsName == "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("This is the default.")
	}
}