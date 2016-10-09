package main

import (
    "fmt"
    "go/types"
)

func main() {
    switch "Marcus" {
    case "Daniel":
        fmt.Println("Wassup Daniel")
    case "Jenny":
        fmt.Println("Wassup Jenny")
    case "Marcus":
        fmt.Println("Wassup Marcus")
        fallthrough
    case "Medhi":
        fmt.Println("Wassup Medhi")
    case "Julian":
        fmt.Println("Wassup Julian")
    case "Sushant":
        fmt.Println("Wassup Sushant")
    }

    multipleEvals()
    noExpression()
    typeSwitchTest()
}

func multipleEvals() {
    switch "Jenny" {
    case "Tim", "Jenny":
        fmt.Println("Wassup Tim, or, err, Jenny")
    case "Marcus", "Medhi":
        fmt.Println("Both of ur names start with M")
    }
}

func noExpression() {
    myFriendsName := "Ma"

    // switch with no expression, executes the the first case that resolves to true, else the default case.

    switch {
    case len(myFriendsName) == 2:
        fmt.Println("Wassup my friend withe name of length of 2")
    case myFriendsName == "Tim":
        fmt.Println("Wassup Tim")
    case myFriendsName == "Jenny":
        fmt.Println("Wassup Jenny")
    case myFriendsName == "Marcus", myFriendsName == "Medhi":
        fmt.Println("Wassup Marcus or Medhi")
    default:
        fmt.Println("nothing matched, this is default.")
    }
}

type Contact struct {
    greeting string
    name string
}

func switchOnType(x interface{}) {
    switch x.(type) {
    case int:
        fmt.Println("int")
    case string:
        fmt.Println("string")
    case Contact:
        fmt.Println("contact")
    default:
        fmt.Println("unknown")
    }
}

func typeSwitchTest() {
    switchOnType(7)
    switchOnType("some string")
    contact := Contact{ "greetings", "Tim" }
    switchOnType(contact)
}