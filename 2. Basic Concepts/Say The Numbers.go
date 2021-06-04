/*
You are making a robot that can speak numbers.
Your robot should take 3 numbers in the range of 0-10 as input and output the corresponding texts in English.

Sample Input:
8
0
5

Sample Output:
Eight
Zero
Five
*/

package main
import "fmt"

func main() {
    var num1, num2, num3 int
    fmt.Scanln(&num1)
    fmt.Scanln(&num2)
    fmt.Scanln(&num3)

    printNum(num1)
    printNum(num2)
    printNum(num3)
}

func printNum(num int){
    if (num <= 10){
      switch num {
      case 0:
        fmt.Println("Zero")
      case 1:
        fmt.Println("One")
      case 2:
        fmt.Println("Two")
      case 3:
        fmt.Println("Three")
      case 4:
        fmt.Println("Four")
      case 5:
        fmt.Println("Five")
      case 6:
        fmt.Println("Six")
      case 7:
        fmt.Println("Seven")
      case 8:
        fmt.Println("Eight")
      case 9:
        fmt.Println("Nine")
      case 10:
       fmt.Println("Ten")
      }
    }
}
