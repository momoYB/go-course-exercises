package main

import "fmt"

func main() {
	type Person struct {
		firstName, lastName string
		favIceCreamFlavour  []string
	}

	me := Person{
		firstName:          "Momchil",
		lastName:           "Borisov",
		favIceCreamFlavour: []string{"Yogurt", "Blueberry"},
	}
	bob := Person{
		firstName:          string("Bob"),
		lastName:           string("Johnson"),
		favIceCreamFlavour: []string{"Chocolate"},
	}

	fmt.Printf("name: %v %v\nfavorite ice cream flavors: ", me.firstName, me.lastName)
	for _, v := range me.favIceCreamFlavour {
		fmt.Printf("%v ", v)
	}
	fmt.Println("\n")

	fmt.Printf("name: %v %v\nfavorite ice cream flavors: ", bob.firstName, bob.lastName)
	for _, v := range bob.favIceCreamFlavour {
		fmt.Printf("%v ", v)
	}
}
