package main

import "fmt"

func main() {
    defer fmt.Println("world") // This execution is postponed
    fmt.Println("hello")       // This runs immediately
}


// func main() {
// 	fmt.Println("counting")

// 	for i := 0; i < 10; i++ {
// 		defer fmt.Println(i)
// 	}

// 	fmt.Println("done")
// }


// When a function returns, its deferred calls are executed in last-in-first-out order.