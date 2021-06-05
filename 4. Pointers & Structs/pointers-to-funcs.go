/*
  The change() function takes an integer argument and changes its value.
  The change_ptr() function does the same using a pointer.

  When you run the code, you will see that the change() function did not change the value of our x variable,
  because the argument is just a copy of its value, while the change_ptr() did change the actual value of x, because it used its memory address as the argument.
  Note that we need to pass the memory address using the & operator to functions that take a pointer as their argument.
*/
package main

import "fmt"

func change(val int) {
  val = 8
}
func change_ptr(ptr *int) {
  *ptr = 8
}

func main() {
  x := 42

  change(x)
  fmt.Println(x) //This will still give 42

  change_ptr(&x)
  fmt.Println(x) //8
}
