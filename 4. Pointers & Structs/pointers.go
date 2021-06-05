//The * operator can also be used to change the value of the memory address the pointer holds:
//We are able to change the value of x using the pointer p.
package main

import "fmt"

func main() {
  x := 42
  p := &x

  *p = 8
  fmt.Println(*p) //8
  fmt.Println(x) //8
}
