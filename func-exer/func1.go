package main

import (
	"fmt"
)

func main() {
	test()
	test1("Hi KD")
	w := woo("HFN")
	fmt.Println(w)
	x, y := mouse("sahaj", "marg")
	fmt.Println(x)
	fmt.Println(y)
}

func test() {
	fmt.Println("func1 called")
}

func test1(s string) {
	fmt.Println(s)
}

func woo(s string) string {
	return fmt.Sprintln("Hello ", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprintln(fn, " -", ln)
	b := true
	return a, b
}
