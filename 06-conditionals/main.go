package main

import "fmt"

func main() {
	x := 15
	y := 15

	// if else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "black"

	if color == "red" {
		fmt.Println("Color is Red")
	} else if color == "blue" {
		fmt.Println("Color is Blue")
	} else {
		fmt.Println("Color is neither Red nor Blue")
	}

	// switch
	switch color {
	case "red":
		fmt.Println("Color is Red")
	case "blue":
		fmt.Println("Color is Blue")
	default:
		fmt.Println("Color is neither Red nor Blue")
	}
}
