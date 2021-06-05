package main

import "fmt"

type Contact struct {
  name string
  age int
}

func (ptr *Contact) increase(val int) {
  ptr.age += val
}

func main() {
  x := Contact{"James", 42}
  x.increase(8)
  fmt.Println(x.age)
}

/*
Now our increase() method uses a pointer as the receiver and is able to modify the age field of the struct it is called on:
Go automatically dereferences the pointers, so we simply call the method with the struct name, just as we did with a simple receiver.
Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
  */
