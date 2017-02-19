package main

import (
    "sort"
    "fmt"
)

func main () {
    s := []string{"Zeno", "Al", "Jen"}
    sort.Strings(s)
    fmt.Println(s)

    //Reverse
    sort.Sort(sort.Reverse(sort.StringSlice(s)))
    fmt.Println(s)
}
