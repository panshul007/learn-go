package main

import "fmt"

func main() {
  age := 44
  fmt.Println(&age) // mem address of age

  changeMe(&age)

  fmt.Println(&age) // same mem address
  fmt.Println(age)  // 24
}

//Chnages the value of age by changing the value stored at its memory location
func changeMe(z *int) {
  fmt.Println(z)  // mem address of age
  fmt.Println(*z) // 44
  *z = 24
  fmt.Println(z)  // same mem address of age
  fmt.Println(*z) //24
}

// in case of slice, string, maps, do not need to pass the address as they are already a referenced type
// everything we pass as a function parameter it is pass by value by default.
