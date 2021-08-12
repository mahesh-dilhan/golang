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

	s = append(s[:2], s[3:]...)
	fmt.Println(s) //deleete index 2

	cntry := GetCountries(2)
	fmt.Println(cntry, cap(cntry))

	accumelate(10, 20, 30, 40, 50, 60)
}

func GetCountries(size int) []string {
	c := [...]string{"USA", "IND", "SL", "RUS"}
	s := make([]string, size)
	copy(s, c[:size])
	return s
}

func accumelate(num ...int) {
	s := []int{}
	s = append(s, num...)
	fmt.Println(s)

}
