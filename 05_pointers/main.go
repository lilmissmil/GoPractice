package main

import "fmt"

func main() {
	a := 42

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	// this is called referencing
	// the above code makes b a pointer to the memory address where an int is stored
	// b is of type "int pointer"
	// *int -- the * is part of the type -- b is of type *int

	fmt.Println(*b)
	// this is called dereferencing
	// because we are asking for the value in the memory address
	// the * is an operator here

	*b = 43        // b says, "The value at this address, change it to 43"
	fmt.Println(a) // 43

	// this is useful because we can pass a memory address instead of a bunch of values
	// (we can pass a reference)
	// and then we can still change the value of whatever is stored at that memory address
	// this makes our programs more performant
	// we don't have to pass around large amounts of data
	// we only have to pass around addresses
	// everything is PASS BY VALUE in go, btw
	// when we pass a memory address, we are passing a value

	fmt.Println(&a)
}
