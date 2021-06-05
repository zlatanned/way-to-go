# FUNCTIONS

## Functions (Code available)

    => Keyword : func

## Arguments (Code available)

    => the argument is provided in the parentheses and includes the name followed by the type.
    => The argument behaves like a variable inside the function's body, meaning you can use its value by its name.
        Eg: Eg: func sum(x, y int) {
                    fmt.Println(x+y)
                }

## Return Statement (Code available)

    => The return statement terminates the function and returns the provided value to the code that called it.
    => Notice that we need to define the return type of the function after the arguments definition. In our case, it is int.
            Eg: func sum(x, y int) int {
                    return x+y
                }
