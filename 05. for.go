package main

import "fmt"

func main() {
	var i = 0

	// while
	for i < 10 {
		fmt.Printf("i: %d ", i)
		i++
	}
	fmt.Println()

	// for
	for j := 0; j < 10; j++ {
		fmt.Printf("j: %d ", j)
	}
	fmt.Println()

	// continue
	for k := 0; k < 10; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Printf("k: %d ", k)
	}
}
