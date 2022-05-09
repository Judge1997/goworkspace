package main

import "fmt"

type dog int

func main() {
	var a dog = 5
	// Type Conversion
	var b = int(a)
	fmt.Printf("The type of variable a is %T, and the value is %v\n", a, a)
	fmt.Printf("The type of variable b is %T, and the value is %v\n", b, b)
}
