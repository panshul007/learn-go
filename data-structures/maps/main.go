package main

import "fmt"

// map is un-ordered
// built up on hash table
// simple key value pairs
// the map is a reference type, i.e, variable stores only the pointer to the address of
// memory location where the data is stored.

func main() {
  m := make(map[string]int) // declaring an empty map using make.
  m["k1"] = 7
  m["k2"] = 13

  delete(m, "k2") // -> on delete the value is set to 0 for this key.
  fmt.Println("map:", m)

  // Print if value is present for the key.
  _, ok := m["k2"]
  fmt.Println("ok?:", ok)

  // Print the value if value is present for the key.
  v, ok := m["k2"]
  fmt.Println("ok?:", ok, v) // if value is missing... it returns 0.

  // Declaring and initializing a map in the same statement.
  n := map[string]int{"foo": 1, "bar": 2}
  fmt.Println("map:", n)
  greeting()
  mapRangeLoop()
}

func greeting() {
  myGreeting := map[int]string{
    0: "Good morning!",
    1: "Bonjour!",
    2: "Buenos dias!",
    3: "Bongiorno!",
  }

  fmt.Println(myGreeting)

  if val, exists := myGreeting[2]; exists {
    delete(myGreeting, 2)
    fmt.Println("val: ", val)
    fmt.Println("exists ", exists)
  } else {
    fmt.Println("That value does not exist.")
    fmt.Println("val: ", val)
    fmt.Println("exits: ", exists)
  }

  fmt.Println(myGreeting)
}

func mapRangeLoop() {
  myGreeting := map[int]string{
    0: "Good morning!",
    1: "Bonjour!",
    2: "Buenos dias!",
    3: "Bongiorno!",
  }

  for key, val := range myGreeting {
    fmt.Println(key, " - ", val)
  }
}
