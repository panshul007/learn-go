package main

import (
  "github.com/panshul007/learn-go/webapp/factory-pattern/appliances"
  "fmt"
)

func main() {
  fmt.Println("Enter preferred appliance type:")
  fmt.Println("0: Stove")
  fmt.Println("1: Fridge")

  var myType int
  fmt.Scan(&myType)

  myAppliance, err := appliances.CreateAppliance(myType)

  if err == nil {
    myAppliance.Start()
    fmt.Println(myAppliance.GetPurpose())
  } else {
    fmt.Println(err)
  }
}