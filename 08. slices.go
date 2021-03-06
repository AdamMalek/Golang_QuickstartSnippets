package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3}
	var b []int = a

	fmt.Println(a, b)

	var c [][]int = [][]int{[]int{4, 5}, []int{6, 7, 8}}

	fmt.Println("fmt.Println")
	fmt.Println(c)

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
