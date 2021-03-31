package main

import "fmt"

type valueReceiverStruct struct{}

func (value valueReceiverStruct) testMethod(message string) {
	fmt.Println("[Value receiver]", message)
}

type pointerReceiverStruct struct{}

func (pointer *pointerReceiverStruct) testMethod(message string) {
	fmt.Println("[Pointer receiver]", message)
}

type testInterface interface {
	testMethod(message string)
}

func main() {
	valueStruct := valueReceiverStruct{}
	pointerStruct := pointerReceiverStruct{}

	valueStruct.testMethod("Direct call from valueStruct")
	pointerStruct.testMethod("Direct call from pointerStruct")

	fmt.Println("We can also have pointers referencing interface implementing instances")
	// we can't get pointer directly
	// var pointerToValueStruct *testInterface = &valueStruct
	// var pointerToPointerStruct *testInterface = &(&pointerStruct)

	valueAsInterface := testInterface(valueStruct)
	pointerAsInterface := testInterface(&pointerStruct)
	var pointerToValueStruct *testInterface = &valueAsInterface
	var pointerToPointerStruct *testInterface = &pointerAsInterface

	// we can't use testMethod from pointer, as it receives value, not pointer
	// pointerToValueStruct.testMethod("Called from pointer")
	// pointerToPointerStruct.testMethod("Called from pointer")
	// we need to dereference this pointer
	(*pointerToValueStruct).testMethod("Called from pointer")
	(*pointerToPointerStruct).testMethod("Called from pointer")

	// In lines above we can see, that we need to manually dereference this pointer to run testMethod.
	// When I was trying to find cause of that, I've found this StackOverflow post: https://stackoverflow.com/a/21283083/8839621
	// TLDR: There is no point in using pointer interfaces and we don't have to worry about copying whole value
	// when passing to function, as interface value takes small amount of memory (2 words)
	// Quote from answer above (author: Volker):
	// "The novice Go programmer just doesn't use pointers to interfaces (as this won't do any good)
	//  and the experienced Go programmer doesn't use pointers to interfaces (as it won't do much good)
	//  unless he needs to modify an interface value, typically during reflection."
}
