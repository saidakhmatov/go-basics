package main

import (
    "fmt"
)
    
func main() {
    for i := 1; i <= 100; i++ {
        FizzBuzz(i)
    }
}

func FizzBuzz(n int) int{
    resp := 0
    if n%15==0{
        fmt.Println("fizzbuzz")
    }else if n%5==0{
        fmt.Print("fizz")
    }else if n%3==0{
        fmt.Println("buzz")
    }else
    {
        fmt.Println(n)
    }
    return resp
}