package main

import "fmt"

func main() {
	var a [3]int = [3]int{1, 2, 3}
	var b [3]int = a

	fmt.Println(a, b)

	var c [3][3]int = [3][3]int{[3]int{4, 5}, [3]int{6, 7, 8}}

	fmt.Println("range")
	for i, t := range c {
		fmt.Printf("Table %d\n", i)
		for _, item := range t {
			fmt.Print(item, " ")
		}
		fmt.Println()
	}

	fmt.Println("fmt.Println")
	fmt.Println(c)

	for i, t := range c {
		fmt.Printf("Table %d\n", i)
		for _, item := range t {
			fmt.Print(item, " ")
		}
		fmt.Println()
	}

	fmt.Println("for len")
	for i := 0; i < len(c); i++ {
		fmt.Printf("Table %d\n", i)
		for j := 0; j < len(c[i]); j++ {
			item := c[i][j]
			fmt.Print(item, " ")
		}
		fmt.Println()
	}
}
