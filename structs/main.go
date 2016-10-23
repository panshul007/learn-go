package main

import "fmt"

// Also known as aggregate type -> aggregating or encapsulating a bunch of fields.
type person struct {
    first string
    last string
    age int
}

type foo int

func main() {
    p1 := person{"James", "Bond", 20}
    p2 := person{"Miss", "Moneypenny", 18}

    foo1 := 42

    fmt.Println(p1.first, p1.last, p1.age)
    fmt.Println(p2.first, p2.last, p2.age)

    fmt.Printf("%T %v \n", foo1, foo1)
}
