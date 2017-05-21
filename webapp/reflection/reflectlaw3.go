package main

import (
  "reflect"
  "fmt"
)

func main() {
  var x1 float32 = 5.8
  v := reflect.ValueOf(x1)
  fmt.Println("v settable? ", v.CanSet()) // false, since it holds a copy of x1. to make it settable v need to hold a reference of x1.

  var x2 float32 = 5.8
  v2 := reflect.ValueOf(&x2)
  fmt.Println("v2 settable? ", v2.CanSet()) // v2 is still not settable. as now v2.CanSet refers to the pointer reference of x2 and not the value stored at the pointer location.

  var x3 float32 = 5.8
  v3 := reflect.ValueOf(&x3)
  fmt.Println("v3 settable? ", v3.CanSet())
  vElem := v3.Elem() // x3
  fmt.Println("vElem settable? ", vElem.CanSet()) // true
  vElem.SetFloat(2.2) // works
  fmt.Println("x3: 5.8 -> ", x3) // 2.2 not 5.8 -> value changed at reference.
}
