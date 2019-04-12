package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 101
}

func bar() (int, string) {
	return 777, "momoYB"
}
