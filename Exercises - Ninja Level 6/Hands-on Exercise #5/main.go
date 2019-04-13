package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	r float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

type shape interface {
	area() float64
}

func info(sh shape) {

	fmt.Println(sh.area())
}

func main() {
	s := square{length: 5}
	c := circle{r: 4}

	info(s)
	info(c)
}
