package main

import "fmt"

func valuecopy(ival int) {
   ival = 0
}

func pointercopy(iptr *int) {
  *iptr = 0
} 

func main() {

 i := 1
 fmt.Println(i)
 valuecopy(i)
 fmt.Println(i)
 fmt.Println(&i)
 pointercopy(&i)
 fmt.Println(i)
 fmt.Println(&i)

}
