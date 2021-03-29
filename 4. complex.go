package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var a = complex128(1i)

	fmt.Println(a)
	fmt.Println(a * a)
	fmt.Println(cmplx.Pow(a, 3))
}
