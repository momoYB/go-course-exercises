package main

import "fmt"

func main() {
	slice2D := [][]string{{"Hello", "World"}, {"Здравей", "Свят"}, {"你好", "世界"}}

	for i := range slice2D {
		fmt.Printf("slice %d:\n", i)
		for _, v2 := range slice2D[i] {
			fmt.Printf("%s ", v2)
		}
		fmt.Println("\n")
	}
}
