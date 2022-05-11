package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Money", "Penny", "Hello, James"}
	xp := [][]string{jb, mp}
	for i, sl := range xp {
		for j, v := range sl {
			fmt.Println(i, j, v)
		}
	}
}
