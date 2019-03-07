package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("This shouldn't print.")
	case true:
		fmt.Println("This should print.")
	}
}
