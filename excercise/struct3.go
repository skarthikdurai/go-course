package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type sedan struct {
	vehicle
	luxury bool
}
type truck struct {
	vehicle
	fourWheel bool
}

func main() {

	sed := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: true,
	}

	tru := truck{
		vehicle: vehicle{
			doors: 5,
			color: "white",
		},
		fourWheel: true,
	}

	fmt.Println(tru.vehicle)
	fmt.Println(sed)

}
