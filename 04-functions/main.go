package main

import "fmt"

// Greeting - Greets the name passed
func Greeting(name string) string {
	return "Hello " + name + "!"
}

// getSum - Outputs the sum of two nos.
func getSum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(Greeting("Sarthak"))
	fmt.Println(getSum(3, 4))
}
