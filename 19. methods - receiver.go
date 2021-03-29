package main

import "fmt"

type letter struct {
	content string
}

func main() {
	myLetter := letter{content: "Sample text"}

	fmt.Println("Original content", myLetter.content)
	myLetter.changeContentValue("Content changed by value")
	fmt.Println("After changeContentValue", myLetter.content)
	myLetter.changeContentPointer("Content changed by pointer")
	fmt.Println("After changeContentPointer", myLetter.content)
	// when invoking method by value object gets copied
	fmt.Println("Original object address", myLetter)
	fmt.Println("Original object address", myLetter.showAddress())
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
