package main

import "fmt"

type person struct {
	first string
	last  string
	age   uint
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}

func main() {
	gosho := person{first: "Georgi", last: "Toshev", age: 27}
	gosho.speak()
}
