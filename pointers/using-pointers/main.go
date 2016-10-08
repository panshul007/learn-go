package main

import "fmt"
// dereferencing

func main() {
    a := 43

    fmt.Println(a)
    fmt.Println(&a)

    var b *int = &a
    // b is a int pointer variable (defined by *int) which stores the memory address of a,
    // or points to the memory address of a.

    fmt.Println(b) // b is referencing the memory address
    fmt.Println(*b) // *b is dereferencing the memory address, so it prints the value of the memory address.

    *b = 42 // make the value at the referenced memory address to 42.
    fmt.Println(a) // should print 42

    // Everything in go is PASS BY VALUE
    // when we pass a memory address, we pass the value.
}
