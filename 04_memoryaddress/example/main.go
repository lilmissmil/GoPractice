package main

import "fmt"

const meterstoyards = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * meterstoyards
	fmt.Println(meters, " meters is ", yards, "yards.")

}
