package main

import "fmt"

type person struct {
	name string
}

func (p *person) SayHi() {
	fmt.Println("Hi, my name is", p.name)
}

func main() {
	adam := person{name: "Adam"}
	adam.SayHi()

	FunctionReceivingObjectOfEmptyInterfaceType(adam)
	FunctionReceivingObjectOfEmptyInterfaceType(1)
	FunctionReceivingObjectOfEmptyInterfaceType(true)
	FunctionReceivingObjectOfEmptyInterfaceType(nil)
	FunctionReceivingObjectOfEmptyInterfaceType("Text")
	FunctionReceivingObjectOfEmptyInterfaceType([]int{1, 2, 3})

	fmt.Println("it might look like interface{} means that any type will match, but it's not true")

	intArray := []int{1, 2, 3}
	// Code below will fail, because []int != []interface{}
	// FunctionReceivingArrayOfEmptyInterfaceTypeObjects(intArray)

	fmt.Println("We need to convert []T to []interface{} to be able to use this object in function below")
	converted := make([]interface{}, len(intArray))
	for i, val := range intArray {
		converted[i] = val
	}
	FunctionReceivingArrayOfEmptyInterfaceTypeObjects(converted)
}

func FunctionReceivingObjectOfEmptyInterfaceType(object interface{}) {
	fmt.Println("1. Function called with argument", object)
}

func FunctionReceivingArrayOfEmptyInterfaceTypeObjects(object []interface{}) {
	fmt.Println("2. Function called with arguments", object)
}
