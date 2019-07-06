package main

import "fmt"

func main() {
	ids := []int{33, 66, 56, 31, 75, 87, 24, 90, 19}

	// Loop through ids using range
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add all the IDs
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum = ", sum)

	// Range with map
	memberID := map[string]string{
		"Bob":    "AGD342",
		"Sharon": "TGM275",
		"Claire": "FPQ390",
	}

	for k, v := range memberID {
		fmt.Printf("%s : %s\n", k, v)
	}

	for k := range memberID {
		fmt.Println("Name: ", k)
	}
}
