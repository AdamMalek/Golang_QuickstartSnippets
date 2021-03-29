package main

import "fmt"

func main() {
	lambdaFunction := func(text string) {
		fmt.Println("Printed from lambda", text)
	}

	lambdaFunction("Text 1")
	lambdaFunction("Text 2")

	functionUsingPassedFunction(15, func(x int) string {
		return fmt.Sprintf(">_%d_<", x) // string interpolation
	})

	// same but with lambda stored to variable for readability
	formatter := func(x int) string {
		return fmt.Sprintf(">_%d_<", x) // string interpolation
	}
	functionUsingPassedFunction(68, formatter)

	func() {
		fmt.Println("Printed from IIFE")
	}()

	multiplyBy2 := functionReturningFunction(2)
	multiplyBy3 := functionReturningFunction(3)

	fmt.Print("15 multiplied by 2 is", multiplyBy2(15))
	fmt.Print("15 multiplied by 3 is", multiplyBy3(15))
}

func functionUsingPassedFunction(number int, formatter func(int) string) {
	fmt.Println("Formatted number:", formatter(number))
}

func functionReturningFunction(multiplier int) func(int) int {
	return func(x int) int {
		return x * multiplier
	}
}
