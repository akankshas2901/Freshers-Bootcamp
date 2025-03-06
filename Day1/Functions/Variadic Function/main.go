package main

import "fmt"

// Function that takes multiple integers as input
func sum(nums ...int) {  
    fmt.Print(nums, " ")  // Print the numbers as a slice
    total := 0            

    // Loop through each number and add it to total
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)  // Print the total sum
}

func main() {
    sum(1, 2)          // Calls sum with two numbers
    sum(1, 2, 3)       // Calls sum with three numbers

    nums := []int{1, 2, 3, 4} 
    sum(nums...)       // Pass a slice to the function
}
