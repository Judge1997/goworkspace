package main

import "fmt"

func main() {
	var a int = 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	a <<= 1
	fmt.Printf("%d\t%b\t%#x", a, a, a)
}
