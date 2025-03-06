// Think of closures as a function inside another function that "remembers" the variables from the outer function—even after the outer function has finished running.

package main

import "fmt"

// This function returns another function (a closure)
func intSeq() func() int {
    i := 0  // i is initialized once

    return func() int { // Anonymous function (closure)
        i++  // Increases i by 1
        return i // Returns updated i
    }
}

func main() {
    nextInt := intSeq()  // nextInt is now a function that remembers `i`

    fmt.Println(nextInt()) // Output: 1
    fmt.Println(nextInt()) // Output: 2
    fmt.Println(nextInt()) // Output: 3

    // Create a new counter (with its own `i` variable)
    newInts := intSeq()
    fmt.Println(newInts()) // Output: 1 (because it's a new instance)
}
