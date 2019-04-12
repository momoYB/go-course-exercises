package main

import "fmt"

func main() {

	defer fmt.Println(
		" A defer statement defers the execution of a function\n",
		"until the surrounding function returns.\n",
		"In this case after the end of the main function.",
	)

	fmt.Println("end of `func main`\n")
}
