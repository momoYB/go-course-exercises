package main

import "fmt"

func main() {
	func() {
		fmt.Println("I am an anonymous func")
	}()
}
