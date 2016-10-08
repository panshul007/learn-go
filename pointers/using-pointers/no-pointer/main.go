package main

import "fmt"

// Passed by value
func zero(z int) {
    fmt.Println("%p\n ", &z) // prints hex value with leading 0x
    fmt.Println(&z) // address in func zero
    z = 0
}

func main() {
    x := 5
    fmt.Println("%p\n ", &x) // prints hex value with leading 0x
    fmt.Println(&x) // address in func zero
    zero(x)
    fmt.Println(x) // x is still 5
}
