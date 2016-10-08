package main

import "fmt"
// referencing

func main() {
    a := 43

    fmt.Println(a)
    fmt.Println(&a)

    var b *int = &a
    // b is a int pointer variable (defined by *int) which stores the memory address of a,
    // or points to the memory address of a.

    fmt.Println(b)
}
