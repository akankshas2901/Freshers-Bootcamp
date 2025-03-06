package main

import (
    "fmt"
    "math"
)

// Define an interface named `Abser`
// Any type that has an `Abs()` method returning a float64 automatically implements this interface
type Abser interface {
    Abs() float64
}

// Define a custom type `MyFloat`
type MyFloat float64

// Implement the `Abs` method for `MyFloat` (making it implement `Abser`)
func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f) // Convert negative value to positive (absolute value)
    }
    return float64(f)
}

// Define a struct `Vertex` to represent a point in 2D space
type Vertex struct {
    X, Y float64
}

// Implement the `Abs` method for *Vertex (pointer receiver)
func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y) // Calculate distance from origin (0,0)
}

func main() {
    var a Abser // Declare a variable of type `Abser` interface

    f := MyFloat(-math.Sqrt2) // Assign a negative float value to `MyFloat`
    v := Vertex{3, 4}         // Create a `Vertex` struct

    a = f  // ✅ `MyFloat` implements `Abser`, so it can be assigned to `a`
    fmt.Println(a.Abs()) // Calls `MyFloat.Abs()`

    a = &v // ✅ `*Vertex` (a pointer to Vertex) implements `Abser`
    fmt.Println(a.Abs()) // Calls `(*Vertex).Abs()`

    // ❌ Error: `v` (a value type) does NOT implement `Abser`
    // Because `Abs()` is defined only on `*Vertex` (pointer type), not `Vertex`
    a = v // Compilation error
}
