package main

import "fmt"

func main() {
	firstVar := 5
	fmt.Printf("dec: %d\tbin: %b\thex: %#x\n", firstVar, firstVar, firstVar)
	secondVar := firstVar << 1
	fmt.Printf("dec: %d\tbin: %b\thex: %#x\n", secondVar, secondVar, secondVar)
}
