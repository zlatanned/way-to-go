# BASIC CONCEPTS

## Variables

    => Variables are used to store values.
    In Go, the var keyword is used to declare variables.
    For example:
        var i int

    The code above declares a variable named i of type int.
    int stands for integer and represents whole numbers.
    
    =>Variable Initialization Methods:
            Way 1 -> var i, j int = 8, 42
            Way 2 -> var i, j = 8, 42
            Way 3 -> i := 8 (Go supports short variable declaration using :=)
            Way 4 -> x, y := 10, 5

## Data Types
    
    => We used int to declare integer values -> var i int = 8 
        Other common types Go supports:
        float32 - a single-precision floating point value.
        float64 - a double-precision floating point value.
        string - a text value.
        bool - Boolean true or false.
        
        Note: The difference between float32 and float64 is the precision, meaning that float64 will represent the number with higher accuracy, but will take more space in memory.

## Constants
    => In some cases your program may need values that are preserved during the program. These are called constants and they cannot be changed from their initial value.
        Constants are declared like variables, but with the const keyword and need to be assigned a value:
            cont pi = 3.14
        Now, pi is a constant and cannot be changed.

        Note: Constants cannot be declared using the := syntax.
        
## Operators - arithmatic, relational and logical (code available)

## Taking Input
    => The "fmt" package also allows you to take input from the user of the program.
        To take input, we need to use the Scanln() function and provide it with the variable which should hold the input value:
            package main
            import "fmt"

            func main() {
              var input string
              fmt.Scanln(&input)
              fmt.Println(input) 
            }
    => Note the ampersand '&' before the variable name -- it is used to return the address of a variable.

## Conditional Statements - if else example added (code available)
    => Note: Go doesn't support ternany operators. You need to use traditional if-else.

## Switch (code available)
    => A switch statement is a shorter way of writing a sequence of if/else statements.
    Note : In Go, the break statement is not needed, as the switch cases evaluate from top to bottom, stopping when a case succeeds.

## Loops (code available)
    => Way 1: for i := 1; i <= 3; i++ { // something }
    => Way 2 : for sum <= 1000 {//something} (similar to while in other languages)
