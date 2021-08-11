package main

import "fmt"

func main() {

	r := []rune("Hello")
	r[0] = 'G'
	buf := string(r)
	fmt.Println(buf)

	s := fmt.Sprintf("%c%c, world!", 72, 'i')
	fmt.Println(s) // "Hi, world!"
}
