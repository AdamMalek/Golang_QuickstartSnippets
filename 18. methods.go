package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	oscar := person{name: "Oscar", age: 95}
	// value got converted to pointer
	oscar.greetByPointer()
	oscar.greetByValue()

	alex := &person{name: "Alex", age: 25}
	alex.greetByPointer()
	// pointer got dereferenced
	alex.greetByValue()
}

func (p *person) greetByPointer() {
	fmt.Printf("[Pointer] Hello, my name is %s and I'm %d years old\n", p.name, p.age)
}

func (p person) greetByValue() {
	fmt.Printf("[Value] Hello, my name is %s and I'm %d years old\n", p.name, p.age)
}
