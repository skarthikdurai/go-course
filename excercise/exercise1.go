package main

import (
	"fmt"
)

var x1 int = 23
var y string = "Heartfulness"
var z bool = true

type hfn int

var a hfn
var b int

func main() {
	x := 42
	a = 1945
	//fmt.Println("Hello, playground")
	//fmt.Println("x =", x)

	//fmt.Println(" values =", x1, y, z)
	//fmt.Printf("%T", x)
	//fmt.Printf("%T\t%T\t%T\t", x, y, z)
	fmt.Printf("%v\t%v\t%v\n", x, y, z)
	s := fmt.Sprintf("%T\t%T\t%T\t", x, y, z)
	fmt.Println(s)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = int(a)
	fmt.Println(b)
	fmt.Printf("%T", b)

	//fun()
	/*
		for i := 0; i < 5; i++ {
			defer fmt.Printf("%d ", i)
		}
	*/
}

func fun() {
	fmt.Println("X1 = ", x1)
}
