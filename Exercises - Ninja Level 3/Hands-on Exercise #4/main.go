package main

import "fmt"

func main() {
	year := 1998
	for {
		fmt.Println(year)
		if year == 2019 {
			break
		}
		year++
	}
}
