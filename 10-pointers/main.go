package main

import "fmt"

func main() {
	a := 5
	b := &a // Assigning 'b' to a pointer of 'a'

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// Use * to read value from the address
	fmt.Println(*b)

	// Change Value with pointer
	*b = 10 // Changes the value of a
	fmt.Println(a)
}
