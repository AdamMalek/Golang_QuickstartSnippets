package main

import "fmt"

func main() {
	generator1 := idGenerator(1)
	generator2 := idGenerator(100)

	printGeneratedItems := func(generator func() int, n int) {
		for i := 0; i < n; i++ {
			fmt.Print(generator(), " ")
		}
		fmt.Println()
	}

	printGeneratedItems(generator1, 5)
	printGeneratedItems(generator2, 5)
}

func idGenerator(startNumber int) func() int {
	currentId := startNumber
	return func() int {
		id := currentId
		currentId++
		return id
	}
}
