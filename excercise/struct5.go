package main

import "fmt"

type chat struct {
	name    string
	message string
	center  string
}

type whatsapp struct {
	msg []chat
}

func main() {

	chat1 := chat{
		name:    "kd",
		message: "test1",
		center:  "chennai",
	}

	chat2 := chat{
		name:    "durai",
		message: "test2",
		center:  "singapore",
	}
	appMsg := whatsapp{
		msg: []chat{
			chat1,
			chat2,
		},
	}
	fmt.Println(appMsg)

}
