package main

import (
    "fmt"
    "math"
)

// Define a struct called `Vertex` to represent a point in 2D space
type Vertex struct {
    X, Y float64
}

// Scale method modifies the original Vertex by multiplying X and Y by `f`
// Uses a pointer receiver (*Vertex) so that changes affect the original object
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

// Abs method calculates the distance of the Vertex from the origin (0,0)
// Uses a pointer receiver, but since it doesn't modify the struct, a value receiver would also work
func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y) // Distance formula: sqrt(X² + Y²)
}

func main() {
    // Create a pointer to a Vertex struct with X=3, Y=4
    v := &Vertex{3, 4}

    // Print the initial values of X, Y and the distance from origin
    fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())

    // Scale the Vertex by 5 (X and Y will be multiplied by 5)
    v.Scale(5)

    // Print the updated values after scaling
    fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
