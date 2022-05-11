package main

import "fmt"

func main() {
	//solution 1
	x := make([]string, 0, 50)
	fmt.Printf("x: %v\n", x)
	fmt.Printf("len of x: %v\n", len(x))
	fmt.Printf("cap of x: %v\n", cap(x))

	x = append(x, "Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming")
	fmt.Printf("x: %v\n", x)
	fmt.Printf("len of x: %v\n", len(x))
	fmt.Printf("cap of x: %v\n", cap(x))

	//solution 2
	y := make([]string, 50)
	fmt.Printf("y: %v\n", y)
	fmt.Printf("len of y: %v\n", len(y))
	fmt.Printf("cap of y: %v\n", cap(y))

	states := []string{"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming"}
	for i := 0; i < len(y); i++ {
		y[i] = states[i]
	}
	fmt.Printf("y: %v\n", y)
	fmt.Printf("len of y: %v\n", len(y))
	fmt.Printf("cap of y: %v\n", cap(y))
}
