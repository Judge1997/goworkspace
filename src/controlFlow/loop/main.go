package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(i, j)
		}
	}

	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	y := 0
	for {
		if y > 5 {
			break
		}

		if y%2 != 0 {
			y++
			continue
		}

		fmt.Println(y)
		y++
	}

	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#U\n", i, i)
	}
}
