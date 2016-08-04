package main

import "fmt"

func max(numbers ...float64) float64 {
	fmt.Printf("%T \n", numbers)
	var largest float64         // initialized to 0
	for _, v := range numbers { //loop or range over it
		if v > largest { // i.e. number greater than zero
			largest = v // then set it equal to the number and loop again
		}
	}
	return largest
}

func main() {
	greatest := max(4, 7, 9, 123, 43, 23, 435, 53, 125)
	fmt.Println(greatest)
}
