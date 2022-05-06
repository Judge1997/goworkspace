package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
	"rsc.io/quote"
)

func Hello() string {
	return quote.Hello()
}

func main() {
	fmt.Println(stringutil.Reverse("Hello"))
	foo()
	fmt.Println("Something more")

	for i := 0; i < 100; i++ {
		if i%9 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("and then we exited")
}

//control flow:
//(1) sqquence
//(2) loop; iterative
//(3) conditional
