package main

import "fmt"

func main() {
	xs := []int{1, 2, 3, 4, 5}
	fmt.Println(foo(xs...))
	fmt.Println(bar(xs))
}

func foo(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func bar(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
