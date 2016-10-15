package main

import "fmt"

func hello() {
    fmt.Print("Hello ")
}

func world() {
    fmt.Println("world")
}

func main() {
    defer world() // the execution of this method is held back till right before the exit of the method it is being called in.
    hello()
}
