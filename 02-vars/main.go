package main

import "fmt"

var name = "Sarthak"

func main() {
	// MAIN TYPES & OTHER INFO
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Sarthak"
	var age int32 = 19
	const isCool = true
	var height float32 = 190.5

	// Shorthand - Must be used only inside a function
	// name := "Sarthak"
	city, email := "Hyderabad", "sarthak@mail.com"

	fmt.Println(name, age, isCool, height, city, email)
	fmt.Printf("%T\n", height)
}
