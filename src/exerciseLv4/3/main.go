package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("x[:5]: %v\n", x[:5])
	fmt.Printf("x[5:]: %v\n", x[5:])
	fmt.Printf("x[2:7]: %v\n", x[2:7])
	fmt.Printf("x[1:6]: %v\n", x[1:6])
}
