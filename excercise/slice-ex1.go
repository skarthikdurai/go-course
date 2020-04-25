package main

import (
	"fmt"
)

func main() {

	x := []int{3, 4, 5, 6, 7}
	fmt.Println(x)

	x = append(x, 32, 44, 300)

	fmt.Println(x)

	y := []int{11, 22, 33}
	fmt.Println(y)

	x = append(x[0:1], x[3:]...)

	fmt.Println(x)
	fmt.Println("===", append(x, y[1:]...))

}
