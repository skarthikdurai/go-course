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

func main() {

	p1 := person{
		first:   "HFN",
		last:    "Heartfulness",
		age:     43,
		country: []string{"India", "USA"},
	}

	p2 := person{
		first:   "Sahaj",
		last:    "Marg",
		age:     1945,
		country: []string{"SG", "JP"},
	}

	m := map[string]person{
		p1.first: p1,
		p2.first: p2,
	}

	for k, v := range m {
		fmt.Println(" = ", k, v.last)
		for _, val := range v.country {
			fmt.Println(val)
		}
	}
}
