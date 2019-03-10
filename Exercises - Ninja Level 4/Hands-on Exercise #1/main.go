package main

import "fmt"

func main() {
	arr := [5]int{0, 1, 2, 3, 4}
	for _, v := range arr {
		fmt.Printf("%v\n", v)
	}
	fmt.Printf("%T", arr)
}
