package main

import "fmt"

func main() {
	if true {
		fmt.Println(1)
	}

	if false {
		fmt.Println(2)
	}

	if !true {
		fmt.Println(3)
	}

	if !false {
		fmt.Println(4)
	}

	if x := 2; x == 2 {
		fmt.Println(5)
	}
	// fmt.Println(x) undefined x because x in if scope

	y := 2
	if y == 1 {
		fmt.Println("y == 1")
	} else if y == 2 {
		fmt.Println("y == 2")
	} else {
		fmt.Println("y not equal 1 and 2")
	}
}
