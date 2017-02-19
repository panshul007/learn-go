package main

import (
    "fmt"
    "math"
)

/**
to implement an interface the type only needs to implement the
 methods signature of the interface.. they do not need to specify 'implements'
 */

type Square struct {
    side float64
}

type Circle struct {
    radius float64
}

func (z Square) area() float64 {
    return z.side * z.side
}

func (c Circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

type Shape interface {
    area() float64
}

func info(z Shape) {
    fmt.Println(z)
    fmt.Println(z.area())
}

func main() {
    s := Square{10}
    c := Circle{20}
    info(s)
    info(c)
}
