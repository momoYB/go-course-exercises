package main

import "fmt"

func main() {
	mySlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	//[0 1 2 3 4]
	slice1 := mySlice[:5]
	fmt.Println(slice1)

	//[5 6 7 8 9]
	slice2 := mySlice[5:]
	fmt.Println(slice2)

	//[2 3 4 5 6]
	slice3 := mySlice[2:7]
	fmt.Println(slice3)

	//[1 2 3 4 5]
	slice4 := mySlice[1:6]
	fmt.Println(slice4)

}
