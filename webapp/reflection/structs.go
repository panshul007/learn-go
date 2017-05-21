package main

import (
  "reflect"
  "fmt"
)

func main() {
  type myStruct struct {
    Field1 int `alias:"f1" desc:"field number 1"` // struct field tags
    Field2 string `alias:"f2" desc:"field number 2"`
    Field3 float64 `alias:"f3" desc:"field number 3"`
  }
  mys := myStruct{2, "Hello", 2.4}
  InspectStructType(&mys)

}

func InspectStructType(i interface{}) {
  mysRValue := reflect.ValueOf(i)
  // to be able to change the values of struct fields, we need to dereference it.
  if mysRValue.Kind() != reflect.Ptr {
    return
  }
  mysRValue = mysRValue.Elem() // value object is changeable now.

  // to make sure we inspect only structs.
  if mysRValue.Kind() != reflect.Struct {
    return
  }

  // changing value ->
  mysRValue.Field(0).SetInt(15)

  mysRType := mysRValue.Type() // reflect.TypeOf(i) // since we can use the actual object now.

  fieldCount := mysRType.NumField() // number fields in the type.

  for i:=0; i<fieldCount; i++ {
    fieldRType := mysRType.Field(i) // type object of each field in mys. // datatype: StructField
    fieldRValue := mysRValue.Field(i) // value object of the field.
    fmt.Printf("Field Name: '%s', field type: '%s', field value: '%v' \n", fieldRType.Name, fieldRType.Type, fieldRValue.Interface())
    fmt.Println("Struct tags, alias: ", fieldRType.Tag.Get("alias"), " decription: ", fieldRType.Tag.Get("desc"))
  }
}
