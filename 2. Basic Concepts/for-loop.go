package main

import "fmt"

func main() {
  sum1 := 0
  sum2 := 1
  res := 0
  //way 1
  for i := 1; i <= 3; i++ {
    sum1 += i
  }
  fmt.Println(sum1)
  
  //way 2
  for sum2 <=3 {
    res += sum2
    sum2++
  }
  fmt.Println(res) 
}
