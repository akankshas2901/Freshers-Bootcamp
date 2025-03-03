package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")

	// Switch statement with a short declaration
	// runtime.GOOS is a built-in function that returns the operating system name.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.") // macOS
	case "linux":
		fmt.Println("Linux.") // Linux
	default:
		// Handles all other OS types like Windows, FreeBSD, etc.
		fmt.Printf("%s.\n", os)
	}
}
