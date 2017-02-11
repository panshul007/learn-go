package main

import "fmt"

// iota -> begins from 1 and increments by 1 after ever use.

const (
  _  = iota             // 0
  KB = 1 << (iota * 10) // 1 << (1 * 10) value of iota -> 1
  MB = 1 << (iota * 10) // 1 << (2 * 10) value of iota -> 2
  GB = 1 << (iota * 10) // 1 << (3 * 10) value of iota -> 3
)

func main() {
  fmt.Println("binary\t\tdecimal")
  fmt.Printf("%b\t", KB)
  fmt.Printf("%d\t\n", KB)
  fmt.Printf("%b\t", MB)
  fmt.Printf("%d\t\n", MB)
  fmt.Printf("%b\t", GB)
  fmt.Printf("%d\t\n", GB)
}
