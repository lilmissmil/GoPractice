package main

import "fmt"

//stick to shorthand or var
//declare, assign and initialize

//var g string
//which means I have a variable identified as b that I am declaring as type string
//hence, declare (name and type)

//to assign, g = "cowboy" (variable is something)

//when you declare and assign at the same time, such as:
//var g string = "cowboy"
//that is called initialize (name and type and variable is something)

//%v is a value in default format
//%T will show the type of the value (i.e. string)

func main() {
	//this is shorthand
	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%v - %T \n", a, a)
	fmt.Printf("%v - %T \n", b, b)
	fmt.Printf("%v - %T \n", c, c)
	fmt.Printf("%v - %T \n", d, d)

	//this is var and prints out zero values
	//we did not assign or initialize it to any value
	//so go will automatically assign a 0 value

	var e int
	var f string
	var g float64
	var h bool

	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Printf("%v \n", h)

}
