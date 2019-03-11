package main

import "fmt"

func main() {
	mySlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	mySlice = append(mySlice, 10)
	fmt.Println(mySlice)

	mySlice = append(mySlice, 11, 12, 13, 14)
	fmt.Println(mySlice)

	s := []int{15, 16, 17, 18, 19}
	mySlice = append(mySlice, s...)
	fmt.Println(mySlice)
}
