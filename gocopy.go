package main

import "fmt"

func main() {
	var src, dst []int
	src = []int{1, 2, 3}
	dst = make([]int, len(src))
	n := copy(dst, src)
	fmt.Println("dst:", dst, "(copied", n, "numbers)")

}
