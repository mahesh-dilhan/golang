package main

import "fmt"

func main() {

 typeOfMe := func(t interface{}) {
  switch i := t.(type) {
   case bool:
    fmt.Println("this is boolean")
   default: 
    fmt.Printf("no idea %T\n", i)
  }

}

 typeOfMe(true)
 typeOfMe(10)
}
