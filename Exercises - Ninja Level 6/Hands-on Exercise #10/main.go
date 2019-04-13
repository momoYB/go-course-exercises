package main

import "fmt"

func multiTen() func() int {
	i := 1
	return func() int {
		i *= 10
		return i
	}
}

func main() {
	num := multiTen()
	for i := 0; i < 3; i++ {
		fmt.Println(num())
	}

	newNum := multiTen()
	fmt.Println(newNum())
}
