package main

import "fmt"

var z = "Bond, James"

// Declare there is a VARIABLE with IDENTIDIFIER "a" and that the VARIABLE with the IDENTIDIFIER "a" is of TYPE int and then ASSIGN the ZERO VALUE of TYPE int to the IDENTIDIFIER "a"
// ZERO VALUE: false for boolean, 0 for integers, 0.0 for floats, "" for string and nil for pointers,  functions, interfaces, slices, channels and maps
var a int

func main() {
	// Short declaration operator
	// Declare a VARIABLE and ASSIGN a VALUE
	x := 42
	fmt.Println(x)
	x = 99
	fmt.Println(x)

	// Full declaration
	var y int = 100 + 70
	fmt.Println(y)

	// Global scope
	fmt.Println(z)
	foo()

	fmt.Println(a)
}

func foo() {
	fmt.Println("Again,", z)
}
