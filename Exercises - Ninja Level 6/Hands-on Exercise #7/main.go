package main

import "fmt"

func main() {
	f := func() { fmt.Println(`Variables can also be of type "func"`) }
	f()
	fmt.Printf("f is of type %T", f)
}
