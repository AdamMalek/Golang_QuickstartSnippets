package main

import "fmt"

// same as in C/C++
func main() {
	var i int = 3

	// 2 ways of defining pointer
	// var ptr *int = &i
	// var ptr3 *int = new(int)
	ptr := &i

	fmt.Println("Initial value", i)
	fmt.Println("Variable address", ptr)
	var pointerToPointer **int = &ptr
	fmt.Println("Pointer address", &ptr)
	fmt.Println("Pointer to pointer value", pointerToPointer)
	fmt.Println("Pointer to pointer value after *", *pointerToPointer)
	fmt.Println("Pointer to pointer value after **", **pointerToPointer)

	modifyByValue(i)
	fmt.Println("After passing variable by value it does not change, because object is copied", i)
	modifyByPointer(ptr)
	fmt.Println("After passing variable by pointer it does change", i)

	function := func() int {
		return 5
	}
	funcPtr := &function
	fmt.Println("We can have pointers to functions as well", &function)
	res := (*funcPtr)()
	fmt.Println("Invocation result", res)

	tab := []int{10, 20, 30}
	ptrTab := &tab
	ptrTab2 := &tab[0]
	fmt.Println(*ptrTab)
	fmt.Println(*ptrTab2)
	// we can't do pointer arithmetics (at least in v1.15)
	// ptr++;
}

func modifyByValue(num int) {
	num += 1
}

func modifyByPointer(num *int) {
	*num += 1
}
