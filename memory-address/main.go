package main

import "fmt"

const metersToYards float64 = 1.09361

// Scan -> always puts the entered values into the memory address.

func main() {
  a := 43
  fmt.Println("a's memory address - ", &a)

  var meters float64
  fmt.Print("Enter meters swam: ")
  fmt.Scan(&meters)
  yards := meters * metersToYards
  fmt.Println(meters, " meters is ", yards, " yards.")
}
