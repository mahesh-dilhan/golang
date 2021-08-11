package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Println("list", arr)

	arr[2] = 10
	fmt.Println("list", arr)

	var a int = 10
	fmt.Println(a)

	for _, value := range arr {
		fmt.Println(value)
	}

	var sl []int = arr[1:2]

	fmt.Println(sl)
}
