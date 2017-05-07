package main

import "fmt"

func main() {
  printType("Text")
  printType(3)
  printType(4.0)
}

//  empty interface means it accepts any type of argument - very useful for writing generic code.
func printType(i interface{}) {
  switch i := i.(type) { // type casts the received value into its inferred type.
  case string:
    fmt.Println("This is a string type", i)
  case int:
    fmt.Println("This is a int type", i)
  case float64:
    fmt.Println("This is a float type", i)
  }
}
