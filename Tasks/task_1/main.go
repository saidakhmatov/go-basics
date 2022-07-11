package main

import "fmt"

func main() {
    var a, b int = 3, 4

    fmt.Printf("a = %d, b = %d\n", a, b)
    //
    b, a = a, b
    //
    fmt.Printf("a = %d, b = %d\n", a, b)
}

