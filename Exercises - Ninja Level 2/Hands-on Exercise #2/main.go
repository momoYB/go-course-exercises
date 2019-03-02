package main

import "fmt"

func main() {
	a := true == true
	b := 62 <= 20
	c := 77 >= 77
	x := "abrakadabra" != "abrakadabra"
	y := -5 < 22
	z := 100-37 > 6*6

	fmt.Printf("a = %v\nb = %v\nc = %v\nx = %v\ny = %v\nz = %v\n", a, b, c, x, y, z)

}
