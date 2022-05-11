package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":  {"Shaken, not stirred", "Martinis", "Woman"},
		"penny_money": {"James Bond", "Literature", "Computer Science"},
		"no_dr":       {"Being evil", "Ice cream", "Sunsets"},
	}
	for k, s := range m {
		fmt.Printf("Key: %v\n", k)
		for i, v := range s {
			fmt.Printf("Index: %v\tValue: %v\n", i, v)
		}
	}

	fmt.Println()

	m["EIEI"] = []string{"Coco", "Stir fried Holy Basil", "Water"}
	for k, s := range m {
		fmt.Printf("Key: %v\n", k)
		for i, v := range s {
			fmt.Printf("Index: %v\tValue: %v\n", i, v)
		}
	}
}
