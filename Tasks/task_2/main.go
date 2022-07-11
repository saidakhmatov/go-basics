package main

import "fmt"

func main() {
    var a, b int = 3, 4

    
    fmt.Printf("a = %d, b = %d\n", a, b)


    fmt.Println(a, "is odd: ", isOdd(a))
    fmt.Println(b, "is even: ", isEven(b))
}

func isEven(num int) bool {
	//
	if num%2==0{
        return true
    }else {
        return false 
    }
	//
}

func isOdd(num int) bool {
	//
    if num%2==0{
        return false
    }else {
        return true 
    }
    //
}