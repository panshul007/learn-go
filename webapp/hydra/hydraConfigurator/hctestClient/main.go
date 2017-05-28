package main

import (
  "github.com/panshul007/learn-go/webapp/hydra/hydraConfigurator"
  "fmt"
)

type ConfS struct {
  TS string `name:"testString" xml:"testString" json:"testString"`
  TB bool `name:"testBool" xml:"testBool" json:"testBool"`
  TF float64 `name:"testFloat" xml:"testFloat" json:"testFloat"`
  TestInt int
}

func main() {
  configstruct := new(ConfS)

  //err := hydraConfigurator.GetConfiguration(hydraConfigurator.CUSTOM, configstruct, "configfile.conf")
  err := hydraConfigurator.GetConfiguration(hydraConfigurator.JSON, configstruct, "configfile.conf")

  if err != nil {
    fmt.Println("Error occured while marshalling: ", err)
  }

  fmt.Println(*configstruct)

  if configstruct.TB {
    fmt.Println("bool is true")
  }

  fmt.Println(float64(4.8 * configstruct.TF))

  fmt.Println(5 * configstruct.TestInt)

  fmt.Println(configstruct.TS)
}
