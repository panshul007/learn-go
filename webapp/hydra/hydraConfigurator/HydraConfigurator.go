package hydraConfigurator

import (
  "errors"
  "reflect"
  "fmt"
)

const (
  CUSTOM uint8 = iota
)

var wrongTypeError error = errors.New("Type must be a pointer to a struct")

func GetConfiguration(confType uint8, obj interface{}, filename string) (err error) {
  // check if this is type pointer
  mysRValue := reflect.ValueOf(obj)
  if mysRValue.Kind() != reflect.Ptr || mysRValue.IsNil() {
    return wrongTypeError
  }

  // get and confirm the struct value
  mysRValue = mysRValue.Elem() // gives the underlying object, not copy of object value.
  if mysRValue.Kind() != reflect.Struct {
    return wrongTypeError
  }

  switch confType {
  case CUSTOM:
    fmt.Println("Marshalling the config file: ", filename)
    err = MarshalCustomConfig(mysRValue, filename)
  }
  return err
}
