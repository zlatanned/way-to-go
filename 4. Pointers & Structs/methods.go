//Both ways are aceeptable. Difference lies in welcome() declaration as a method with receiver struct arg or a normal function and in its calling respectively.

package main

import "fmt"

type Contact struct {
  name string
  age int
}

func welcome(x Contact) {
  fmt.Println(x.name)
  fmt.Println(x.age)
}
func main() {
  x := Contact{"James", 42}
  welcome(x)
}

// package main

// import "fmt"

// type Contact struct {
//   name string
//   age int
// }

// func (x Contact) welcome() {
//   fmt.Println(x.name)
//   fmt.Println(x.age)
// }

// func main() {
//   x := Contact{"James", 42}
//   x.welcome()
// }
