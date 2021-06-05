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
                
##  Multiple Return (Code available)
    => func swap(x, y int) (int, int) {
            return y, x
        }

## Defer (Code available)
    => A defer statement ensures that the function is called only after the surrounding function returns.
    => Note: defer is often used for cleanup, for example, to release resources used by the code, such as files, connections, etc.
    
## Multiple Defer (Code available)
    => If you have deferred multiple function calls, they will execute in last-in-first-out order. 
    The defer calls are stacked on top of each other, which is why they are executed in last-in-first-out order.
    
## Scope (Code available)
    => Scope is where a variable can be used.
        There are two main scopes in Go: local and global.
        A variable defined in the function is called a local variable. Their scope is only in the function body, which means they only exist within their function.
        Note:Global variables are often considered a bad practice. It is better to pass variables as function arguments.
