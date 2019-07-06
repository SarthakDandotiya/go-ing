package main

import "fmt"

func main() {
	// Arrays
	var fruits [2]string

	// Assign Values
	fruits[0] = "Apple"
	fruits[1] = "Orange"

	// Declare and Assign
	citys := [2]string{"Delhi", "Mumbai"}

	fmt.Println(fruits)
	fmt.Println(fruits[1])
	fmt.Println(citys)

	// Slices
	names := []string{"Jack", "Jill", "Bill", "Will", "Jane", "Bob"}
	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(names[1:4])
}
