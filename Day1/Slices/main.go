package main

import (
    "fmt"
    "slices" // Importing slices package for utility functions
)

func main() {
    // 1️⃣ Declaring an uninitialized slice (nil slice)
    var s []string
    fmt.Println("uninit:", s, s == nil, len(s) == 0)
    // Output: uninit: [] true true
    // - `s` is nil, meaning it has not been allocated memory
    // - `len(s) == 0` confirms it has zero elements

    // 2️⃣ Creating a slice with `make`
    s = make([]string, 3) // Creates a slice of length 3, initialized with empty strings
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
    // Output: emp: [  ] len: 3 cap: 3

    // 3️⃣ Setting and getting values
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)  // Output: set: [a b c]
    fmt.Println("get:", s[2]) // Output: get: c

    // 4️⃣ Getting length
    fmt.Println("len:", len(s)) // Output: len: 3

    // 5️⃣ Appending elements (adding more values)
    s = append(s, "d")      // Adding one element
    s = append(s, "e", "f") // Adding multiple elements
    fmt.Println("apd:", s)  // Output: apd: [a b c d e f]

    // 6️⃣ Copying a slice
    c := make([]string, len(s)) // Create a new slice with the same length as `s`
    copy(c, s)                  // Copy contents of `s` to `c`
    fmt.Println("cpy:", c)      // Output: cpy: [a b c d e f]

    // 7️⃣ Slicing (Extracting sub-slices)
    l := s[2:5]  // Extracts elements from index 2 to 4 (excluding index 5)
    fmt.Println("sl1:", l) // Output: sl1: [c d e]

    l = s[:5]  // Extracts elements from the beginning up to index 4
    fmt.Println("sl2:", l) // Output: sl2: [a b c d e]

    l = s[2:]  // Extracts elements from index 2 to the end
    fmt.Println("sl3:", l) // Output: sl3: [c d e f]

    // 8️⃣ Declaring and initializing a slice in a single line
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t) // Output: dcl: [g h i]

    // 9️⃣ Comparing slices using slices.Equal
    t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
        fmt.Println("t == t2") // Output: t == t2
    }

    // 🔟 Multi-dimensional slices (Slice of slices)
    twoD := make([][]int, 3) // Creating a slice with 3 inner slices
    for i := 0; i < 3; i++ {
        innerLen := i + 1            // Inner slices have increasing lengths
        twoD[i] = make([]int, innerLen) // Allocating space for each inner slice
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j // Assigning values to the inner slices
        }
    }
    fmt.Println("2d: ", twoD) // Output: 2d: [[0] [1 2] [2 3 4]]

    // Another way to declare a multi-dimensional slice
    twoD = [][]int{
        {1, 2, 3},
        {4, 5, 6},
    }
    fmt.Println("2d: ", twoD) // Output: 2d: [[1 2 3] [4 5 6]]
}


// Slices are more powerful than arrays because they can grow dynamically.
// make([]type, length, capacity) is used to create a new slice.
// Appending (append(slice, values...)) increases the size of a slice.
// Copying (copy(dest, src)) duplicates slice contents.
// Slicing (slice[start:end]) extracts a part of a slice.
// Multi-dimensional slices are flexible, unlike arrays where all inner arrays must have the same size.
// Comparing slices can be done using slices.Equal(slice1, slice2).