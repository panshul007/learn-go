package main

import (
  "time"
  "math/rand"
  "errors"
  "fmt"
)

var scMapping = map[string]int {
  "James": 5,
  "Kevin": 10,
  "Rahul": 9,
}

type findError struct {
  Name, Server, Msg string
}

func (e findError) Error() string {
  return e.Msg
}

// Error variable
var ErrCrewNotFound = errors.New("Crew member not found.")

func findSC(name, server string) (int, error) {
  //Simulating searching
  time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)

  if v, ok := scMapping[name]; !ok {
    // return -1, errors.New("Crew member not found.") // Returning a simple new error with message

    // return -1, ErrCrewNotFound // Returning a simple error variable

    // return -1, fmt.Errorf("Crew member %s could not be found on server '%s' ", name, server)

    // return -1, findError{name, server, "Crew member not found."}

      panic("Crew member not found") // not recommended but useful to handle them as many packages use them.
  } else {
    return v, nil
  }
}

func main() {
  rand.Seed(time.Now().UnixNano())
  clearance, err := findSC("Ruko", "Server 1")
  fmt.Println(" Clearance level found: ", clearance, " Error code: ", err)
}

// Checking for existence of error
func main2() {
  rand.Seed(time.Now().UnixNano())
  if clearance, err := findSC("Ruko", "Server 1"); err != nil {
    fmt.Println("Error occured while searching for clearance level: ", err)
  } else {
    fmt.Println(" Clearance level found: ", clearance)
  }
}

// Checking for error type
func main3() {
  rand.Seed(time.Now().UnixNano())
  clearance, err := findSC("Ruko", "Server 1")
  if err == ErrCrewNotFound {
    //handle the error
    fmt.Println(" Confirmed error is crew not found ", err, clearance)
  }
}

// Using fmt.errorf
func main4() {
  rand.Seed(time.Now().UnixNano())
  if _, err := findSC("Ruko", "Server 1"); err != nil {
    fmt.Println("Error Occured: ", err)
  }
}

// Using custom error type, checking error type and getting values out of custom error type by type casting.
func main5() {
  rand.Seed(time.Now().UnixNano())
  if clearance, err := findSC("Ruko", "Server 1"); err != nil {
    fmt.Println("Error occured while searching: ", err)
    if v, ok := err.(findError); ok { // type assertion to check if the error is of the expected type. v stores the type casted value.
      fmt.Println("Server name: ", v.Server)
      fmt.Println("Crew member name: ", v.Name)
    }
  } else {
    fmt.Println("Crew member has security clearance: ", clearance)
  }
}

// recovering from panic
func main6() {
  rand.Seed(time.Now().UnixNano())

  defer func() {
    if err := recover(); err != nil { // not recommended but useful to handle them as many packages use them.
      // logic to recover from panic, to prevent application crash
      fmt.Println("A panic recovered", err)
    }
  }()

  if clearance, err := findSC("Ruko", "Server 1"); err != nil {
    fmt.Println("Error occured while searching: ", err)
    if v, ok := err.(findError); ok { // type assertion to check if the error is of the expected type. v stores the type casted value.
      fmt.Println("Server name: ", v.Server)
      fmt.Println("Crew member name: ", v.Name)
    }
  } else {
    fmt.Println("Crew member has security clearance: ", clearance)
  }
}