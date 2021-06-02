# GETTING STARTED

## Packages

    Every Go program is made up of packages.

    A Go program starts running in the main package.

    This is why we need to declare our code as the main package


## Imports

    => Apart from main, Go has many packages that can be imported and used in the code to accomplish different tasks.
    One of the most popular packages is "fmt", which stands for format, and provides input and output functionality.
    To import a package, we use the import statement:
        
        For single import: import "fmt" 
        For single or multiple : import ( "fmt" "shahi" "zlatan")

    => Each package has exported names, which you can access after importing them.
    In Go, a name is exported if it begins with a capital letter.
    You can access the exported names using the package name, a dot, and the exported name.
    For example, we access the Println() function of the "fmt" package
    
    => Similar to other programming languages such as Java or C++, func main() is the entry point of our program.
    It is the function that gets executed when we run our program.

## Comments
    For single line comment: //This is a comment
    For multi-line comments: 
        /*
            This is an example program
            that only generates a simple
            text output using the "fmt" package
            loving you is
        */
