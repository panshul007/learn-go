package main

import "fmt"

func main() {
  if true {
    fmt.Println("This ran")
  }

  if false {
    fmt.Println("This did not run")
  }

  if !true {
    fmt.Println("This does not run")
  }

  b := true

  if b {
    fmt.Println("food")
  }

  if food := "Chocolate"; b {
    fmt.Println(food)
  }

  if false {
    fmt.Println("no dice")
  } else {
    fmt.Println("else works")
  }

  if false {
    fmt.Println("no dice")
  } else if true {
    fmt.Println("else works")
  } else {
    fmt.Println("bla")
  }
}
