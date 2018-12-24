package main

import "fmt"

type myType int

var x myType
var y int

func main() {

	fmt.Printf("%v %T\n", x, x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Printf("%v %T", y, y)

}
