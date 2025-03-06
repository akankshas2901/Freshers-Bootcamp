package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// Declare `v` inside the if statement
	if v := math.Pow(x, n); v < lim {
		return v  // If v is less than lim, return v
	} else {
		// v is also accessible inside the else block
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// v is NOT accessible here
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10), // 3^2 = 9 (less than 10, so return 9)
		pow(3, 3, 20), // 3^3 = 27 (greater than 20, print message, return 20)
	)
}
