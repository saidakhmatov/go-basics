package main

import (
    "fmt"
    "math"
)


func main() {
    
    fmt.Println("Please input first number:")
    var num1 int
    fmt.Scanln(&num1)
    fmt.Println("Please input second number")
    var num2 int
    fmt.Scanln(&num2)

    DisplayMinimumNumberFunction(num1,num2)
}
func DisplayMinimumNumberFunction(num1 int, num2 int) {
    
    
    res_3 := math.Min(float64(num1), float64(num2))
   
    fmt.Println("The smallest number is: ",res_3)

}