//We could use (*p).age to access the age field of the struct, but that looks complicated and hard to read. Go allows you to shorten that syntax and simply use p.age instead.

package main

import "fmt"

type Contact struct {
  name string
  age int
}

func main() {
  x := Contact{"James", 42}
  p := &x

  fmt.Println(p.age) //42
}
