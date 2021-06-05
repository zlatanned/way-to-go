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
        
## Passing Pointers to Functions (Code avaiable)

## Structs

    => Go does not support classes. Instead, it has structs. Structs are collections of fields that allow you to group data together.
        For example, let's make a struct to store the data of Contacts:
                type Contact struct {
                  name string
                  age  int
                }
        Our Contact struct has two fields, a string and an integer.

        Now, we can create a new Contact using the following syntax:
            x := Contact{"James", 42}
        x is now a structs that is initialized with the data provided in the curly braces.
        We can also provide the names of the fields when creating a new struct.
            For example:
            x := Contact{name: "James", age: 42}
        This makes it easier to read the code.
        
 ## Pointers to Structs (Code Available)
    => Similar to simple pointers, we can also make pointers to structs using the & operator:
        x := Contact{"James", 42}
        p := &x
        Pointers to structs are automatically dereferenced, meaning we can access the field values by simply using a dot:
        We could use (*p).age to access the age field of the struct, but that looks complicated and hard to read.
        Go allows you to shorten that syntax and simply use p.age instead.
