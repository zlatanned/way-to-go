//Pointers to structs are automatically dereferenced, meaning we can access the field values by simply using a dot
//We could use (*p).age to access the age field of the struct, but that looks complicated and hard to read. Go allows you to shorten that syntax and simply use p.age instead.

package main

import "fmt"

type Contact struct {
  name string
  age int
}

func main() {
  x := Contact{"James", 42}
  // or x := &Contact{"James", 42} is also allowed, and then we need to print x.age and not p.age. Also next line wont be required
  p := &x

  fmt.Println(p.age) //42
}
