package main

import "fmt"

type Contact struct {
  name string
  age int
}

func main() {
  x := Contact{"James", 42}

  x.name = "Zlatan"
  x.age = 8
  fmt.Println(x.age) //8
  fmt.Println(x.name) //Zlatan
}
