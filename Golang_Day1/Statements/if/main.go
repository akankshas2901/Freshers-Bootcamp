package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"		//imaginary no
	}
	return fmt.Sprint(math.Sqrt(x))
}
// fmt.Sprint() converts values into a string and returns it without printing it to the console.

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
