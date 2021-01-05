package main

import "fmt"

func main(){
   nxtSeq := nextInt()
   fmt.Println(nxtSeq())
   fmt.Println(nxtSeq())
}

func nextInt() func() int {

i := 0 
return func() int {
 i++
 return i
 
}



}
