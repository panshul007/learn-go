package main

import "fmt"

func main() {
  var x [58] string
  fmt.Println(x)
  fmt.Println(len(x))
  fmt.Println(x[42])

  for i := 65; i <= 122; i++ {
    x[i-65] = string(i)
  }

  fmt.Println(x)
  fmt.Println(len(x))
  fmt.Println(x[42])
  example2()
}

func example2() {
  var x [256] int

  fmt.Println(x)
  fmt.Println(len(x))
  fmt.Println(x[42])

  for i := 0; i < 256; i++ {
    x[i] = i
  }

  // item or index, value
  for i, v := range x {
    fmt.Printf("%v - %T - %b\n", v, v, v)
    if i > 50 {
      break
    }
  }
  //Todo: why do we need i in the second loop?
}
