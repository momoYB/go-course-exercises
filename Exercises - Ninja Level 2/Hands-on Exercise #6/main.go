package main

import "fmt"

const (
	y1 = 2019 - iota
	y2
	y3
	y4
)

func main() {
	fmt.Printf("%v %v %v %v", y1, y2, y3, y4)
}
