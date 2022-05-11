package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Not print1")
	case 2 == 4:
		fmt.Println("Not print2")
	case 3 == 3:
		fmt.Println("Print3")
	case 4 == 4:
		fmt.Println("Print4, but not print")
	}

	fmt.Println()

	switch {
	case false:
		fmt.Println("Not print1")
	case 2 == 4:
		fmt.Println("Not print2")
	case 3 == 3:
		fmt.Println("Print3")
		fallthrough
	case 4 == 4:
		fmt.Println("Print4,but print with fallthrough")
	}

	fmt.Println()

	switch {
	case false:
		fmt.Println("Not print1")
	case 2 == 4:
		fmt.Println("Not print2")
	case 3 == 3:
		fmt.Println("Print3")
		fallthrough
	case 4 == 4:
		fmt.Println("Print4,but print with fallthrough")
		fallthrough
	case 7 == 9:
		fmt.Println("Not true1,but print with fallthrough")
		fallthrough
	case 11 == 14:
		fmt.Println("Not true2,but print with fallthrough")
		fallthrough
	case 14 == 14:
		fmt.Println("true 14,but print with fallthrough")
	}

	fmt.Println()

	switch {
	case false:
		fmt.Println("Not print1")
	case 2 == 4:
		fmt.Println("Not print2")
	default:
		fmt.Println("print")
	}

	fmt.Println()

	switch "Coco" {
	case "a":
		fmt.Println("a")
	case "b":
		fmt.Println("b")
	case "Coco":
		fmt.Println("Coco")
	default:
		fmt.Println("default")
	}

	fmt.Println()

	switch "Coco" {
	case "a", "Coco":
		fmt.Println("a, Coco")
	case "b":
		fmt.Println("b")
	case "c":
		fmt.Println("c")
	default:
		fmt.Println("default")
	}
}
