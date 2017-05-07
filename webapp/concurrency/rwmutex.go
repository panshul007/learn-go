package main

import (
  "sync"
  "time"
  "math/rand"
  "fmt"
)

// creating a thread safe map

type MapCounter struct {
  m map[int]int
  sync.RWMutex
}

func main() {
  mc := MapCounter{m: make(map[int]int)}
  go runWriters(&mc, 10)
  go runReaders(&mc, 10, "reader1")
  go runReaders(&mc, 10, "reader2")
  time.Sleep(15 * time.Second)
}

func runWriters(mc *MapCounter, n int) {
  for i := 0; i<n; i++ {
    mc.Lock()
    println("writing ", i)
    mc.m[i] = i * 10
    mc.Unlock()
    time.Sleep(1 * time.Second)
  }
}

func runReaders(mc *MapCounter, n int, name string) {
  for {
    mc.RLock()
    v := mc.m[rand.Intn(n)]
    mc.RUnlock()
    fmt.Print(name + " ")
    fmt.Println(v)
    time.Sleep(1 * time.Second)
  }
}
