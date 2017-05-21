package main

import "fmt"

func main() {
  X := 5
  ChangeX(&X) //*int => X
  fmt.Println(X)  // value changes to 10
}

func ChangeX(X *int) {
  *X = 10  // pointer de-referenced and new pointer assigned.
}
