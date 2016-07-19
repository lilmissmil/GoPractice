package main

import "fmt"

// example of a nested loop

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, " - ", j)
		}
	}

	// other ways to use the for statement

	h := 0
	for h < 10 {
		fmt.Println(h)
		h++
	}

	// break examples below

	k := 0
	for {
		fmt.Println(k)
		if k >= 10 {
			break
		}
		k++
	}

	m := 0
	for {
		m++
		if m%2 == 0 {
			continue
			// i.e. 3%2 = 1, not 0, so it is false and gets out of the loop, 2%2 = 0 so it is true,
			// which means it's only going to stop this iteration of the loop and go back to the top.
		}
		fmt.Println(m)
		if m >= 50 {
			break
		}
	}
}
