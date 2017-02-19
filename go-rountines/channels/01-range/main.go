package main

import (
    "fmt"
)

// channels allow communication between two go routines.

func main() {
    c := make(chan int)

    go func() {
        for i := 0; i < 10; i++ {
            c <- i
        }
        close(c) // after close... no more values can be put onto the channel. Although values put on the channel
        // before close can still be received.
    }()

    for n := range c {
        fmt.Println(n)
    }
}
