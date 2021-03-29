package main

import "fmt"

func main() {
	// foreach
	iterable := []string{"one", "two", "three"}
	for index, value := range iterable {
		fmt.Println(index, value)
	}

	mapObj := map[string]int{"one": 1, "eleven": 11, "two": 2}
	for key, value := range mapObj {
		fmt.Println(key, value)
	}

	fmt.Println("Using indices only")
	for indexOnly, _ := range []int{1, 2, 3} {
		fmt.Println(indexOnly)
	}

	fmt.Println("Using values only")
	for _, valueOnly := range []int{10, 20, 30} {
		fmt.Println(valueOnly)
	}
}
