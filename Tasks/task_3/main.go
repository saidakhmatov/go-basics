package main

import (
    "fmt"
    "math"
)
    
func main() {
    var r float32 = 10.04

    // fmt.Scanf("%f", &r)
    fmt.Println("R =", r)

    fmt.Printf("Area: %0.2f\n", area(r))
}

func area(r float32) (area float32) {
    //
    circleArea := math.Pi*r*r
    sqrArea := 4*r*r
    return sqrArea-circleArea
    //
}