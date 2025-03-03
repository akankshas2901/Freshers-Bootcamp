package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}

// op: 0 false false false
// A var statement can be at package or function level


// var i, j int = 1, 2

// func main() {
// 	var c, python, java = true, false, "no!"
// 	fmt.Println(i, j, c, python, java)
// }

//op: 1 2 true false no!
// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.


//short hand declarations works only inside a function
// func main() {
// 	var i, j int = 1, 2
// 	k := 3
// 	c, python, java := true, false, "no!"

// 	fmt.Println(i, j, k, c, python, java)
// }


// Go has various value types including strings, integers, floats, booleans, etc. 