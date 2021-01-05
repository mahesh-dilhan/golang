package main 

import "fmt"

func main() {

  var lvalu = "abc"
  var rvalu = "Abc"
  var rs = true
  var decvalu string = "abc"
  

  if(lvalu==rvalu) {
     rs = true
  } else {
    rs = false

  }
  fmt.Println(rs)

  if(lvalu==decvalu) {
   fmt.Println(" no diff dec or default type ")
  }

}

