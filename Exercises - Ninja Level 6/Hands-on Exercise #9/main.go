package main

import "fmt"

func swap(x, y *int) { *x, *y = *y, *x }

func decr(x, y int) bool {
	if x < y {
		return true
	}
	return false
}

func incr(x, y int) bool {
	if x > y {
		return true
	}
	return false
}

func bubbleSort(xs []int, cmpr func(x, y int) bool) {
	for i := 0; i < len(xs)-1; i++ {
		for j := 0; j < len(xs)-i-1; j++ {
			if cmpr(xs[j], xs[j+1]) {
				swap(&xs[j], &xs[j+1])
			}
		}
	}
}
func main() {
	xs := []int{7, 10, 2, 13, 0, -26, 5, 7}
	fmt.Println("unsorted:", xs)
	bubbleSort(xs, incr)
	fmt.Println("sorted in increasing order:", xs)
	bubbleSort(xs, decr)
	fmt.Println("sorted in decreasing order:", xs)

}
