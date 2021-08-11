package main

import "fmt"

func main() {

	m := make(map[int]int)

	for i := 1; i <= 10; i++ {
		m[i] = i
	}

	fmt.Println(m)

	delete(m, 2)

	fmt.Println(m)

	for ID, VALUE := range m {
		fmt.Println(ID, VALUE)
	}

	value, ok := m[1]
	fmt.Println(ok)
	if !ok {
		panic("no value found for key 1")
	}
	fmt.Println(value)
}
