# Arrays, Range, Map

  ## Arrays
    => An array is a sequence of elements of the same type.
    => An array is defined using square brackets which define the number of elements the array will hold.

    For example:
        var a [5] int
        Now a is an array of 5 integers.

    You can also define and initialize values of the array using the following syntax:
      a := [5]int{0, 2, 4, 6, 8}
    Note: As you can see, we need to provide the size of the array when declaring it. This means that arrays have a fixed size.
    
  ## Slices
    An array has a fixed size, meaning once defined, you cannot change the number of elements it holds.
    To overcome this, Go provides the slice, which is a dynamically-sized view into the elements of an array.

    A slice is based on an array and is defined by specifying two indices, a low and high bound, separated by a colon:
        a := [5]int{0, 2, 4, 6, 8}
        var s []int = a[1:3]
    
    This code selects the elements with index 1 to 3 from the a array, including the first given index, but excluding the last.
    So, the slice s will now include the values [2, 4]:
    
    Note: You can omit the low or the high bound.
    Omitting the low bound will take the value 0, while omitting the high bound will take the length of the array.
    For example: a[:3] will take the first 3 elements of the array.
    
  ## make()
    Go provides a make() function to create slices. This is how you create dynamically-sized arrays.
        For example => a := make([]int, 5)
    The make function creates an array of the given type and size, and returns a slice that refers to that array.

    After creating a slice, we can add append new elements to it using the append() function:
    The append() function takes the slice as its first argument and the elements to be added to the end of the slice as its next argument.
    It then returns a new slice, containing the old slice plus the new elements appended.
    You can append multiple values at once by just comma separating the values as arguments.for example: append(s, 1, 2, 3)
