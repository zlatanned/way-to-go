package main

import "fmt"

func main() {
  fmt.Println("start")

  for i := 0; i < 5; i++ {
    defer fmt.Println(i)
  }
  fmt.Println("end")
}

/*
  Output : 
  start
  end
  4
  3
  2
  1
  0
*/
