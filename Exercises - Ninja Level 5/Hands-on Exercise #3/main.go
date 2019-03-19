package main

import "fmt"

type Vehicle struct {
	doors uint
	color string
}

type Truck struct {
	Vehicle
	fourWheel bool
}
type Sedan struct {
	Vehicle
	luxury bool
}

func main() {

	myTruck := Truck{
		Vehicle: Vehicle{
			doors: 5,
			color: "Blue",
		},
		fourWheel: true,
	}

	mySedan := Sedan{
		Vehicle: Vehicle{
			doors: 4,
			color: "Silver",
		},
		luxury: false,
	}

	fmt.Printf("type: %T\ndoors: %d\ncolor: %s\nfour wheel: %v\n\n",
		myTruck, myTruck.doors, myTruck.color, myTruck.fourWheel)

	fmt.Printf("type: %T\ndoors: %d\ncolor: %s\nluxury: %v\n\n",
		mySedan, mySedan.doors, mySedan.color, mySedan.luxury)
}
