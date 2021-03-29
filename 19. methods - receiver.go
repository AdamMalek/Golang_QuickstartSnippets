package main

import "fmt"

type letter struct {
	content string
}

func main() {
	myLetter := letter{content: "Sample text"}

	fmt.Println("Original content", myLetter.content)
	myLetter.changeContentValue("Content changed by value")
	fmt.Println("After changeContentValue:", myLetter.content)
	myLetter.changeContentPointer("Content changed by pointer")
	fmt.Println("After changeContentPointer:", myLetter.content)

	fmt.Println("When invoking method by value object gets copied")
	fmt.Printf("Original object address %p\n", &myLetter)
	fmt.Printf("Copied object address %p\n", myLetter.showAddress())

	// When to use pointer receiver?
	// 1. when we want to modify original objects
	// 2. when struct is big, so we can skip copying large object

	// Golang takes care of converting object to pointer

}

func (l letter) changeContentValue(newContent string) {
	l.content = newContent
}

func (pL *letter) changeContentPointer(newContent string) {
	pL.content = newContent
}

func (l letter) showAddress() *letter {
	return &l
}
