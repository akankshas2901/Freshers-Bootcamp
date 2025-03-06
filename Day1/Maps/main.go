package main

import (
    "fmt"
    "maps" // Importing a package for additional map functions
)

func main() {
    //  Creating an Empty Map
    m := make(map[string]int) 
    // This creates an empty map where:
    // - The key is of type `string`
    // - The value is of type `int`

    //  Adding Key-Value Pairs
    m["k1"] = 7
    m["k2"] = 13
    fmt.Println("map:", m) // Output: map: map[k1:7 k2:13]

    //  Accessing Values Using Keys
    v1 := m["k1"]  // Fetch the value for key "k1"
    fmt.Println("v1:", v1) // Output: v1: 7

    //  Accessing a Non-Existent Key
    v3 := m["k3"] // "k3" does not exist, so it returns 0 (default int value)
    fmt.Println("v3:", v3) // Output: v3: 0

    //  Finding the Number of Key-Value Pairs
    fmt.Println("len:", len(m)) // Output: len: 2

    //  Deleting a Key-Value Pair
    delete(m, "k2") // Removes "k2" from the map
    fmt.Println("map:", m) // Output: map: map[k1:7]

    //  Removing All Key-Value Pairs
    clear(m) // Clears the entire map
    fmt.Println("map:", m) // Output: map: map[]

    //  Checking If a Key Exists
    _, prs := m["k2"] // The second return value tells if the key exists
    fmt.Println("prs:", prs) // Output: prs: false (because "k2" was deleted)

    // Creating and Initializing a Map in One Line
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n) // Output: map: map[foo:1 bar:2]

    //  Comparing Two Maps
    n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2") // Output: n == n2 (since both maps have the same data)
    }
}
