package main

import "fmt"

func split(sum int) (x, y int) { // Named return values: x, y
    x = sum * 4 / 9  // x is assigned a value
    y = sum - x      // y is assigned the remaining value
    return           // No need to specify x, y; they are returned automatically
}

func main() {
    fmt.Println(split(17)) // Calls split(17) and prints the returned values
}


