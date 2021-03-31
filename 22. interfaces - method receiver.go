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

	fmt.Println("Interface 'testInterface' requires, that method called 'testMethod' should exist")
	// we have no problem with 'valueReceiverStruct', because we have this struct defined and we have
	// method named 'testMethod' receiving object of type 'valueReceiverStruct'
	var asInterfaceValue testInterface = valueStruct
	asInterfaceValue.testMethod("Called from reference")

	// here we have a problem, because we have struct 'pointerReceiverStruct' defined, but
	// there is no method called 'testMethod' with 'pointerReceiverStruct' as receiver.
	// We have method with this name, hovever instead of using value it uses pointer '*pointerReceiverStruct' as receiver.
	// Consequence of this is that 'pointerReceiverStruct' does not implement 'testInterface',
	// but '*pointerReceiverStruct' does.
	// that is why we cannot use pointerStruct to assign to 'asInterfacePointer' interface this way:
	// var asInterfacePointer testInterface = pointerStruct
	// and we have to assign '*pointerReceiverStruct', which we can get by using:
	var asInterfacePointer testInterface = &pointerStruct
	asInterfacePointer.testMethod("Called from reference")

	// same problem will occur, when we'll try to populate []interface:
	// - valueReceiverStruct implements testInterface, so we can provide 'valueStruct' directly
	// - pointerReceiverStruct does not implement testInterface, but *pointerReceiverStruct does,
	// so we have to convert pointerStruct to *pointerReceiverStruct by getting it's address: &pointerStruct
	var interfaceImplementations []testInterface = []testInterface{valueStruct, &pointerStruct}
	for _, obj := range interfaceImplementations {
		obj.testMethod("Called from interfaceImplementations array")
	}

	fmt.Println("We can also have pointers referencing interface implementing instances")
	// we can't get pointer directly
	// var pointerToValueStruct *testInterface = &testInterface(valueStruct)
	// var pointerToPointerStruct *testInterface = &(&pointerStruct)

	var pointerToValueStruct *testInterface = &asInterfaceValue
	var pointerToPointerStruct *testInterface = &asInterfaceValue

	// we can't use testMethod from pointer, as it receives value, not pointer
	// pointerToValueStruct.testMethod("Called from pointer")
	// pointerToPointerStruct.testMethod("Called from pointer")
	// we need to dereference this pointer
	(*pointerToValueStruct).testMethod("Called from pointer")
	(*pointerToPointerStruct).testMethod("Called from pointer")
}
