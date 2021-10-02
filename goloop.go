package main

import "fmt"

func main() {

	i := 1
	for i <= 10 {
		fmt.Printf("this is %d ", i)
		i++
	}

	fmt.Println("")
	sum := 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 { // skip odd numbers
			continue
		}
		sum += i
	}
	fmt.Println(sum) // 6 (2+4)

}
