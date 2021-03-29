package main

import "fmt"

func main() {
	var a = make([]int, 3)
	fmt.Println(a)

	a[1] = 10

	a = append(a, 1)
	fmt.Println(a)

	a = append(a, 2, 3)
	fmt.Println(a)

	b := a[2:5] // get slice of a for indices 2 to 4
	fmt.Println(b)

	fmt.Println("Slices by range are not copies, but references")
	b[1] = 17
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("copy copies values from src to dest slice, not creating new slice")
	copy(a[2:5], b) //basically copying to same place
	b[1] = 55
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("To get new copy not connected to original slice we need to create empty new one and only then copy")
	b = make([]int, 3)
	copy(a[2:5], b)
	b[1] = 66
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("Partial slices")
	fmt.Println("Original slice", a)

	sliceFromTo := a[1:4]
	fmt.Println("sliceFromTo [1:4]", sliceFromTo)

	sliceFrom := a[1:]
	fmt.Println("sliceFrom [1:]", sliceFrom)

	sliceTo := a[:4]
	fmt.Println("sliceTo [:4]", sliceTo)

	var inline []string
	fmt.Println("inline defined", inline)

	var inline2 []string = []string{"one", "two"}
	fmt.Println("inline2 defined with items", inline2)

	fmt.Println("append returns new slice, old is untouched", inline)
	withAppendedItem := append(inline, "appended")
	fmt.Println(inline, withAppendedItem)
}
