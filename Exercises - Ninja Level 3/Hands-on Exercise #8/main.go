package main

import "fmt"

func main() {
	house := "House Arryn"

	switch house {
	case "House Stark":
		fmt.Println("Winter is Coming")
	case "House Lannister":
		fmt.Println("Hear Me Roar")
	case "House Targaryen":
		fmt.Println("Fire and Blood")
	case "House Baratheon":
		fmt.Println("Ours is the Fury")
	case "House Arryn":
		fmt.Println("As High as Honor")
	case "House Tyrell":
		fmt.Println("Growing Strong")
	case "House Martell":
		fmt.Println("Unbowed, Unbent, Unbroken")
	case "House Tully":
		fmt.Println("Family, Duty, Honor")
	case "House Greyjoy":
		fmt.Println("We Do Not Sow")
	default:
		fmt.Println("Game of Thrones")
	}
}
