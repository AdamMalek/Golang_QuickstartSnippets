package main

import "fmt"

func main() {
	FunctionWithoutReturn(4)
	res := SumNumbersUpTo(15)
	fmt.Println("Sum up to 15 is", res)

	pow2, pow3, pow4, pow5 := GetPowers(4)
	fmt.Println("Powers of 4:", pow2, pow3, pow4, pow5)
}

func FunctionWithoutReturn(a int) {
	fmt.Println("Numers from 0 to", a)
	for i := 0; i <= a; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func SumNumbersUpTo(maxNumber int) int {
	sum := 0
	for i := 0; i <= maxNumber; i++ {
		sum += i
	}
	return sum
}

// multiple return values (tuple)
func GetPowers(numberToPower int) (int, int, int, int) {
	pow2 := numberToPower * numberToPower
	pow3 := pow2 * numberToPower
	pow4 := pow3 * numberToPower
	pow5 := pow4 * numberToPower
	return pow2, pow3, pow4, pow5
}

// if more params have same type it's enough to just specify type on last param of given type
func MultiplyNumbers(a, b, c int, str1, str2 string) int {
	_ = str1
	_ = str2
	return a * b * c
}
