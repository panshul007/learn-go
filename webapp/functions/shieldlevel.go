package main

import "fmt"

type level int

func main() {
  sl := new(level)
  sl.raiseShieldLevel(4)
  sl.raiseShieldLevel(5)

  fmt.Println(*sl)
}

// functions can be attached to any type.
func (lv *level) raiseShieldLevel(i int) {
  if *lv == 0 {
    *lv = 1
  }

  *lv = (*lv) * level(i) // type casting int into level to multiply with level.
}
