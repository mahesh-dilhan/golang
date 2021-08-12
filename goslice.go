package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	for a, b := range s {
		fmt.Println(a, b)
	}

	s2 := []int{10, 11}
	s = append(s, s2...)

	fmt.Println(s)

}
