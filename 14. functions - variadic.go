package main

import "fmt"

func main() {
	VariadicFunction(1, 2, 3)
	VariadicFunction(1, 2, 3, 4, 5, 6, 7, 8)
	array := []int{1, 2, 3}

	// split operator
	VariadicFunction(array...)

}

// similar to JS, or C# params
func VariadicFunction(numbers ...int) {
	fmt.Println("Called with args", numbers)
	for _, val := range numbers {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

// only last argument can be this way
func Variadic2(text string, isTrue bool, numbers ...int) {

}
