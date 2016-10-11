package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
    greet("Jane")
    greet("John")
    greet2("John", "Doe")
    fmt.Println(greet3("Jane", "Doe"))
    fmt.Println(greet4("Jane", "Doe"))
    fmt.Println(greet5("Jane", "Doe"))
    fmt.Println(average(43, 56, 87, 12, 45, 57))

    data := []float64{43, 56, 87,12, 45, 57} //slice of float 64
    n := average(data...) // -> Variadic arguments, take data and expand every element in the slice.
    fmt.Println(n)
    expression()
    wrapExec()
}

func greet(name string) {
    fmt.Println(name)
}

func greet2(fname, lname string) {
    fmt.Println(fname, lname)
}

// prints to a string
func greet3(fname, lname string) string {
    return fmt.Sprint(fname, lname)
}

// named return
func greet4(fname, lname string) (s string) {
    s = fmt.Sprint(fname, lname)
    return
}

// multiple return
func greet5(fname, lname string) (string, string) {
    return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}

// Variadic parameter function
func average(sf ...float64) float64 {
    total := 0.0
    for _, v:= range sf {
        total += v
    }
    return total / float64(len(sf))
}

// func expression

func expression () {
    greeting := func() {
        fmt.Println("Hello World!")
    }

    greeting()
    fmt.Printf("%T\n", greeting) // type -> func()

    greet := makeGreeter()
    fmt.Println(greet())
    fmt.Printf("%T\n", greet) // type -> func() string

}


func makeGreeter() func() string {
    return func() string {
        return "Hello World!"
    }
}

// Closure

func wrapper() func() int {
    x := 0
    return func() int {
        x++
        return x
    }
}

func wrapExec() {
    increment := wrapper()
    fmt.Println(increment())
    fmt.Println(increment())
}