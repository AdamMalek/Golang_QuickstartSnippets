package main

import "fmt"

type country struct {
	name       string
	cities     []string
	getTaxRate func() int
}

func main() {
	// allocating memory manually
	country1 := new(country)
	// pointers are automatically dereferenced, no need to use -> as in C/C++
	country1.name = "Poland"
	country1.cities = []string{"Cracow", "Warsaw", "Chrzęszczyżewoszyce"}
	country1.getTaxRate = func() int { return 23 }
	fmt.Println(*country1)
	fmt.Println(country1.name, "tax rate is", country1.getTaxRate())

	// we can get pointer
	country2 := &country{name: "UK", cities: []string{"London", "Manchester"}}
	fmt.Println("pointer", country2)
	fmt.Println("value", *country2)

	// struct initialization
	country3 := country{name: "Germany"}
	fmt.Println(country3)
}
