# POINTERS & STRUCTS

## Pointers
    => All of the values that we define in our program are stored in the computer memory and have their own unique memory address.
       Pointers are special variables that hold the memory address of values.

       In Go, we declare a pointer using a * (aka dereferencing operator): var p *int
       Now, p is a pointer to an integer value.
       Pointers allow you to pass references to values in your program.
    => We know how to define a pointer, but how do we assign it a memory address? This is done using the & operator, which returns the memory address of a variable.
        For example:
          x := 42
          p := &x
        Now p is a pointer and holds the memory address of x.
        Let's output its value: fmt.Println(p) //0xc0000b8020
        If we want to access the underlying value of a pointer, we can use the * operator : fmt.Println(*p) //42
        The * operator can also be used to change the value of the memory address the pointer holds: We were able to change the value of x using the pointer p. (CODE AVAILABLE)
