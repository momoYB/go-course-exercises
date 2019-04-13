package main

import "fmt"

func greeter() func(s string) string {
	return func(s string) string {
		return "Hello " + s
	}
}

func main() {
	zdr := greeter()
	fmt.Println(zdr("MoMo"))
}
