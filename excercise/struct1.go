package main

import (
	"fmt"
)

type person struct {
	first   string
	last    string
	age     int
	country []string
}

type order struct {
	person
	first   string
	isOrder bool
}

func main() {

	p1 := person{
		first:   "HFN",
		last:    "SM",
		age:     43,
		country: []string{"India", "USA"},
	}

	o1 := order{
		person: person{
			first:   "Sahaj",
			last:    "Marg",
			age:     1945,
			country: []string{"India1", "USA1"},
		},
		isOrder: true,
		first:   "Manapakkam",
	}

	fmt.Println(p1)
	fmt.Println(p1.first)
	fmt.Println(p1.age)

	//fmt.Println(o1)
	fmt.Println(o1.first, o1.person.first, o1.last, o1.age, o1.isOrder)

	for i, v := range p1.country {
		fmt.Println("Country = ", i, v)
	}
}
