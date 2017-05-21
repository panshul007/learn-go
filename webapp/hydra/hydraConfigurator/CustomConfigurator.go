package hydraConfigurator

import (
  "reflect"
  "strconv"
  "errors"
  "os"
  "bufio"
  "fmt"
  "strings"
)

type ConfigFields map[string]reflect.Value

func (f ConfigFields) Add(name, value, t string) error {
  switch t {
  case "STRING":
    f[name] = reflect.ValueOf(value)
  case "INTEGER":
    i, err := strconv.Atoi(value) // Atoi -> ascii to int
    if err != nil {
      return err
    }
    f[name] = reflect.ValueOf(i)
  case "FLOAT":
    fl, err := strconv.ParseFloat(value, 64)
    if err != nil {
      return err
    }
    f[name] = reflect.ValueOf(fl)
  case "BOOL":
    b, err := strconv.ParseBool(value)
    if err != nil {
      return err
    }
    f[name] = reflect.ValueOf(b)
  }
  return nil
}

func MarshalCustomConfig(v reflect.Value, filename string) error {
  fmt.Println("Beginning the mashalling of: ", filename)
  // to recover from any panics.
  defer func() {
    if r := recover(); r != nil {
      fmt.Println("Panic occured: ", r)
    }
  }()

  if !v.CanSet() {
    return errors.New("Value passed is not Settable.")
  }

  fmt.Println("Can set check passed.")
  file, err := os.Open(filename)
  if err != nil {
    return err
  }
  fmt.Println("File opened.")
  defer file.Close()
  fields := make(ConfigFields) // make(map[string]reflect.Value)
  fmt.Println("Beginning to scan the config file: ", filename)
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    fmt.Println("Processing line: ", line)

    // Sanity check : to prevent strings.Split to throw a panic.
    if strings.Count(line, "|") != 1 || strings.Count(line, ";") != 1 {
      continue
    }

    args := strings.Split(line, "|")
    valueType := strings.Split(args[1], ";")
    name, value, vtype := strings.TrimSpace(args[0]), strings.TrimSpace(valueType[0]), strings.ToUpper(valueType[1])
    fields.Add(name, value, vtype)
  }

  // check if any errors occurred while scanning.
  if err := scanner.Err(); err != nil {
    return err
  }

  vt := v.Type()
  for i :=0; i<v.NumField(); i++ {
    fieldType := vt.Field(i)
    fieldValue := v.Field(i)

    name := fieldType.Tag.Get("name")
    if name == "" {
      name = fieldType.Name
    }

    if v, ok := fields[name]; ok {
      fieldValue.Set(v)
    }
  }

  return nil
}