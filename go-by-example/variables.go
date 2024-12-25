package main

import "fmt"

func main() {

    var a = "Initial text"
    fmt.Println(a)

    var b, c int = 100, 202
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e1 int
		var e2 string
		var e3 float32
    fmt.Println(e1)
    fmt.Println(e2)
    fmt.Println(e3)

    f := "Car" // Walrus operator for declaring & initializing a variable. Syntax only available inside functions
    fmt.Println(f)
}