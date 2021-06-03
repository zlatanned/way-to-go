package main

import "fmt"

func main() {
  x := 42
  y := 8
    
  // logical AND
  fmt.Println(x!=y && x<=y)
    
  // logical OR
  fmt.Println(x!=y || x<=y)
    
  // logical NOT
  fmt.Println(!(x>y))
}
