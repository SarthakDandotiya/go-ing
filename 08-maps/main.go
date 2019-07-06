package main

import "fmt"

func main() {
	// Define Map
	emails := make(map[string]string)
	memberID := map[string]string{
		"Bob":    "AGD342",
		"Sharon": "TGM275",
		"Claire": "FPQ390",
	}

	// Assign Keys & Values
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@yahoo.com"
	emails["Claire"] = "claire@gmail.com"
	emails["Mike"] = "claire@rediff.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Sharon"])

	// Delete from map
	delete(emails, "Mike")
	fmt.Println(emails)

	fmt.Println(memberID)
}
