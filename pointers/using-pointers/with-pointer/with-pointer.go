package main

import "fmt"

// Passed by reference
func zero(z *int) {
  fmt.Println(z) // address in func zero
  *z = 0         // de-referencing
}

func main() {
  x := 5
  fmt.Println(&x) // address in func zero
  zero(&x)
  fmt.Println(x) // x is 0, as the memory address value is changed in function zero.
}
