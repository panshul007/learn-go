package main

import "fmt"

type Person struct {
  Name  string
  Age   int
}

type DoubleZero struct {
  Person
  LicenseToKill bool
}

func (p Person) Greeting() {
  fmt.Println("I am just a regular person.")
}

func (dz DoubleZero) Greeting() {
  fmt.Println("Miss Moneypenny, so good to see you.")
}

func main() {
  p1 := Person{
    Name: "James Bond",
    Age: 44,
  }

  p2 := DoubleZero{
    Person: Person{
      Name: "Miss Moneypenny",
      Age:   23,
    },
    LicenseToKill: true,
  }

  // All the fields of embedded type gets promoted to root level.
  p1.Greeting()
  p2.Greeting()
  p2.Person.Greeting()
}

