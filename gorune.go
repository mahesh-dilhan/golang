package main

import "fmt"

func main() {

	r := []rune("Hello")
	r[0] = 'G'
	s := string(r)
	fmt.Println(s)
}
