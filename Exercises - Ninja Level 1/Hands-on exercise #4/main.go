package main

import "fmt"

type myType int

var x myType

func main() {

	fmt.Printf("%v %T\n", x, x)
	x = 42
	fmt.Println(x)

}
