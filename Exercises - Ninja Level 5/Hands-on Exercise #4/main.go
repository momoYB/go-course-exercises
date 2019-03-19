package main

import "fmt"

func main() {
	myAnonymousStruct := struct {
		fieldInt    int
		fieldString string
		fieldBool   bool
		fieldSlice  []uint
		fieldMap    map[uint]string
	}{
		fieldInt:    1,
		fieldString: "one",
		fieldBool:   true,
		fieldSlice:  []uint{1, 2, 3},
		fieldMap: map[uint]string{
			1: "one",
			2: "two",
			3: "three",
		},
	}

	fmt.Println(myAnonymousStruct, "\n")

	for _, v := range myAnonymousStruct.fieldSlice {
		fmt.Println(v)
	}
	fmt.Println()

	for k, v := range myAnonymousStruct.fieldMap {
		fmt.Println(k, " : ", v)
	}
}
