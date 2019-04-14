package main

import "fmt"

type person struct {
	first   string
	last    string
	bastard bool
}

func changeMe(p *person) {
	p.first = "Aegon"
	p.last = "Targaryen"
	p.bastard = false
}

func main() {
	jon := person{
		first:   "Jon",
		last:    "Snow",
		bastard: true,
	}
	fmt.Println(jon)

	changeMe(&jon)
	fmt.Println(jon)
}
