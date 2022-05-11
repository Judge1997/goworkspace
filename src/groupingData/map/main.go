package main

import "fmt"

func main() {
	m := map[string]int{
		"James": 32,
		"Honey": 27,
	}
	fmt.Printf("m: %v\n", m)
	fmt.Printf("m[James]: %v\n", m["James"])
	fmt.Printf("m[eiei]: %v\n", m["eiei"])

	v, ok := m["eiei"]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)

	if v, ok := m["James"]; ok {
		fmt.Printf("m[James]: %v\n", v)
	}

	if v, ok := m["EIEI"]; ok {
		fmt.Printf("m[EIEI]: %v\n", v)
	} else {
		fmt.Printf("Not have m[EIEI] in m\n")
	}
	m["EIEI"] = 50
	fmt.Printf("Added m[EIEI]\n")
	if v, ok := m["EIEI"]; ok {
		fmt.Printf("m[EIEI]: %v\n", v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "EIEI")
	fmt.Printf("Deleted m[EIEI]\n")
	if v, ok := m["EIEI"]; ok {
		fmt.Printf("m[EIEI]: %v\n", v)
	} else {
		fmt.Printf("Not have m[EIEI] in m\n")
	}

	fmt.Printf("m: %v\n", m)
}
