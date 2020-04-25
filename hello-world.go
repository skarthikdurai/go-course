package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello GO First")
	fmt.Println(n)
	fmt.Println(err)

	for i := 0; i < 100; i++ {
		go fmt.Println(i)
	}
}
