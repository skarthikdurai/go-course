package main

import "fmt"

func main() {

	s := struct {
		name    string
		message map[uint16]string
		center  []string
	}{
		name: "kd",
		message: map[uint16]string{
			1: "hello",
			2: "how are you",
		},
		center: []string{
			"chennai",
			"singapore",
		},
	}
	fmt.Println(s)
	fmt.Println(s.center)
	fmt.Println(s.message)

	for _, v := range s.message {
		fmt.Println(v)
	}
}
