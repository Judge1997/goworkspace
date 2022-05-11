package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 20, 4}
	fmt.Printf("x: %v\n", x)
	fmt.Printf("len x: %v\n", len(x))
	fmt.Printf("cap x: %v\n", cap(x))

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Println(x[1])
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	x = append(x, 5, 6, 7)
	fmt.Printf("x: %v\n", x)
	fmt.Printf("len x: %v\n", len(x))
	fmt.Printf("cap x: %v\n", cap(x))

	y := []int{10, 11, 12}
	x = append(x, y...)
	fmt.Printf("x: %v\n", x)
	fmt.Printf("len x: %v\n", len(x))
	fmt.Printf("cap x: %v\n", cap(x))

	x = append(x[:2], x[4:]...)
	fmt.Printf("x: %v\n", x)
	fmt.Printf("len x: %v\n", len(x))
	fmt.Printf("cap x: %v\n", cap(x))

	z := make([]int, 10, 100)
	fmt.Printf("z: %v\n", z)
	fmt.Printf("len z: %v\n", len(z))
	fmt.Printf("cap z: %v\n", cap(z))

	z[0] = 1
	z[9] = 10
	fmt.Printf("z: %v\n", z)
	fmt.Printf("len z: %v\n", len(z))
	fmt.Printf("cap z: %v\n", cap(z))

	z = append(z, 11)
	fmt.Printf("z: %v\n", z)
	fmt.Printf("len z: %v\n", len(z))
	fmt.Printf("cap z: %v\n", cap(z))

	jb := []string{"James", "Bond"}
	fmt.Printf("jb: %v\n", jb)
	mp := []string{"Money", "Penny"}
	fmt.Printf("mp: %v\n", mp)
	xp := [][]string{jb, mp}
	fmt.Printf("xp: %v\n", xp)
}
