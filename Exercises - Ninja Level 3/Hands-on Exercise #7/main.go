package main

import "fmt"

func main() {
	phrase := "Where are my dragons?"
	if phrase == "Valar morghulis!" {
		fmt.Println("Valar dohaeris!")
	} else if phrase == "Chaos isn't a pit." {
		fmt.Println("Chaos is a ladder.")
	} else {
		fmt.Println("Drakaris!")
	}
}
