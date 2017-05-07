package main

import (
  "sync"
  "time"
  "math/rand"
  "fmt"
)

func main() {
  var wg sync.WaitGroup

  for i:=0;i<=5;i++ {
    wg.Add(1)

    go func(i int) {
      defer wg.Done()
      time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
      fmt.Println("Work done for: ", i)
    }(i)
  }

  wg.Wait()
}
