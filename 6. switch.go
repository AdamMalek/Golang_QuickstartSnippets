package main

import "fmt"

func main() {
	var a = 14

	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4:
		fmt.Println("three or four")
	default:
		fmt.Println("other")
	}

	valueForSwitchWithoutExpression := 18
	switch {
	case valueForSwitchWithoutExpression == 18:
		fmt.Println("valueForSwitchWithoutExpression Is 18")
	case a > 5:
		fmt.Println("valueForSwitchWithoutExpression is not 18 and a is greater than 5")
	default:
		fmt.Println("Default case")
	}

	function := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("It's number")
		case bool:
			fmt.Println("It's bool")
		case rune:
			fmt.Println("It's rune xD")
		case string:
			fmt.Println("It's string")
		case float32:
			fmt.Println("It's float32")
		case float64:
			fmt.Println("It's float64")
		case complex64:
			fmt.Println("It's complex")
		}
	}
	function(1)
	function(false)
	function('a')
	function("a")
	function(float32(0.1))
	function(float64(.1))
	function(complex64(1 + 3i))
}
