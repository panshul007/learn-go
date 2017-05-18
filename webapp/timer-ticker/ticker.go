package main

import (
  "time"
  "fmt"
)

func main1() {
  go tickCounter(1)
  time.Sleep(5 * time.Second)
}

func tickCounter(n int) {
  ticker := time.NewTicker(time.Duration(n) * time.Second)
  i := 0
  for t := range ticker.C {
    i ++
    fmt.Println("Count ", i, " at ", t)
  }
}

func main2() {
  ticker := time.NewTicker(1 * time.Second)
  go tickCounter2(ticker)
  time.Sleep(5 * time.Second)
  ticker.Stop() // be careful -> with this command, ticker.C channel does not close properly. need to handle it separately.
  time.Sleep(10 * time.Second)
  fmt.Println("Exiting...")
}

func tickCounter2(ticker *time.Ticker) {
  i := 0
  for t := range ticker.C {
    i ++
    fmt.Println("Count ", i, " at ", t)
  }
}

// to handle proper exit of ticker channel
func main() {
  ticker := time.NewTicker(1 * time.Second)
  done := make(chan bool)
  go tickCounter3(ticker, done)
  time.Sleep(5 * time.Second)
  ticker.Stop() // be careful -> with this command, ticker.C channel does not close properly. need to handle it separately.
  done <- true
  time.Sleep(10 * time.Second)
  fmt.Println("Exiting...")
}

func tickCounter3(ticker *time.Ticker, done chan bool) {
  i := 0
  Loop:
  for {
    select {
    case t := <- ticker.C:
      i++
      fmt.Println("Count ", i, " at ", t)
    case <-done:
      fmt.Println("done signal")
      break Loop
    }
  }
  fmt.Println("Exiting the tick counter...")
}