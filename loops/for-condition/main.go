package main

import "fmt"

func main() {
    i := 0
    for i < 10 {
        fmt. Println(i)
        i++
    }

    // another loop

    j := 0
    for {
        fmt.Println(j)
        if j >= 10 {
            break
        }
        j++
    }

    // yet another using continue.

    k := 0
    for {
        k++
        if k%2 == 0 {
            continue
        }
        fmt.Println(k) // prints only odd numbers
        if k >= 50 {
            break
        }
    }

    // printing Runes

    for l := 0; l <= 140; l++ {
        fmt.Printf("%v - %v - %v \n", l, string(l), []byte(string(l)))
    } // Prints integer, string representation of integer, byte value of string.
}
