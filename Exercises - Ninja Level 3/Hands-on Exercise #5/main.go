package main

import "fmt"

func main() {
	for i := 10; i < 101; i++ {
		fmt.Printf("When% d is divided by 4, the remainder is %d\n", i, i%4)
	}
}
