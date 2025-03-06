package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//random integer between 0 and 9.
	fmt.Println("My favorite number is", rand.Intn(10))
}


// When importing a package, you can refer only to its exported names.Any "unexported" names are not accessible from outside the package.

// Go has a rule:
// If a function, variable, or constant starts with a lowercase letter, it is private (hidden) outside its package.
// If it starts with a capital letter, it is public and can be used in other files.
// ✔ Names starting with a capital letter (like Pi) can be used outside their package.
// ❌ Names starting with a lowercase letter (like pi) are private and hidden.