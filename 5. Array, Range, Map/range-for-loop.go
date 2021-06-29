package main

import "fmt"

func main() {
  a := make([]int, 5)
  a[1] = 2
  a[2] = 3

  for i, v := range a {
    fmt.Println(i, v)
  }
}

/*
  Output:
      0 0
      1 2
      2 3
      3 0
      4 0
*/
