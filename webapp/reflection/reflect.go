package main

import (
  "reflect"
  "fmt"
)

func main() {
  var x1 float32 = 5.7

  inspectIfTypeFloat(x1)

  /*
  type myFloat float64
  var x3 myFloat = 5.7
  v := reflect.ValueOf(x3)
  fmt.Println(v.Type())
  */
}

func inspectIfTypeFloat(i interface{}) {
  // Object to reflection object - Law of reflection 1
  v := reflect.ValueOf(i)
  fmt.Println("type: ", v.Type()) // same as reflect.TypeOf(x1)
  fmt.Println("Is type float64?", v.Kind() == reflect.Float64)
  fmt.Println("Float Value: ", v.Float()) // v.Float returns type float64

  // reflection object to original object - Law of reflection 2
  interfaceValue := v.Interface()
  switch t := interfaceValue.(type) {
  case float32:
    fmt.Println("Original float32 value: ", t)
  }
}
