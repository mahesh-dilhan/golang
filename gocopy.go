package main

import "fmt"

func main() {
	var src, dst []int
	src = []int{1, 2, 3}
	dst = make([]int, len(src))
	n := copy(dst, src)
	fmt.Println("dst:", dst, "(copied", n, "numbers)")

	appendMe(src, dst)

	appendslice()
}

func appendMe(src, dst []int) {
	dst = append(dst, src...)
	fmt.Println("dst:", dst)
}

func appendslice() {
	a := []byte("ba")

	a1 := append(a, 'd')
	a2 := append(a, 'g')

	fmt.Println(string(a1)) // bag
	fmt.Println(string(a2)) // bag
}
