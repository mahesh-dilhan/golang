// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, 世界")
	marray := [3]int32{1, 2, 3}
	fmt.Println(marray[0], marray[1])
	num, deno, err := division(1, 1)
	fmt.Println(num, deno, err)
	var slice1 []int32 = make([]int32, 3, 4)
	fmt.Println(slice1)
	slice1[0] = 1
	slice1 = append(slice1, 10)
	fmt.Println(slice1)
	var mmap map[string]int = map[string]int{"ABC": 10, "CD": 2}
	var b, ok = mmap["B"]
	fmt.Println(b, ok)
	fmt.Printf("%c", "HELLO"[1])
	fmt.Println()
	for k, v := range mmap {
		fmt.Println(k, v, len(k))
		fmt.Printf("%c", k[1])
		fmt.Println()

	}
	var sub = []string{"s", "b"}
	fmt.Println(sub)
	var strw strings.Builder
	for i := range sub {
		strw.WriteString(sub[i])
	}
	fmt.Printf("\n%v", strw.String)
}

func division(nume int, deno int) (int, int, error) {
	return nume / deno, nume % deno, nil
}
