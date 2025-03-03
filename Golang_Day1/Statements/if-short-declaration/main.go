package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// Short variable declaration inside if statement
	if v := math.Pow(x, n); v < lim {
		return v  // If v is less than lim, return v
	}
	return lim  // Otherwise, return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10), // 3^2 = 9 (less than 10, so return 9)
		pow(3, 3, 20), // 3^3 = 27 (greater than 20, so return 20)
	)
}
