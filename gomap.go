package main

import "fmt"


func main() {

 m := make(map[int]int)

for i := 1; i <= 10; i++ {
   m[i] = i
 }
 
 fmt.Println(m)

 delete(m,2)

 fmt.Println(m)

}
