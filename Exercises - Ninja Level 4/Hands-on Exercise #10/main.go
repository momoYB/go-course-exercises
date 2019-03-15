package main

import "fmt"

func main() {
	myMap := map[string][]string{
		"Jon Snow":           {"Knowing nothing.", "Bending the knee.", "Wildlings."},
		"Daenerys Targaryen": {"Breaking the wheel.", "Riding dragons.", "Being called \"Queen\"."},
		"Tyrion Lannister":   {"Drinking and knowing things.", "Trial by combat.", "Accepting what he is and wearing it like armor."},
	}

	myMap["Petyr Baelish"] = []string{"Chaos.", "Money", "Power."}

	delete(myMap, "Petyr Baelish")

	for k, slice := range myMap {
		fmt.Println(k)
		for i, v := range slice {
			fmt.Printf("%d: %v\n", i+1, v)
		}
		fmt.Println()
	}
}
