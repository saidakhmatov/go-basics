package main

import (
	"fmt"

	"github.com/saidakhmatov/golang/Project_5/bigint"
)

func main() {

	// declaring number a
	a, err := bigint.NewInt("+++0+00000-2034032")
	if err != nil {
		panic(err)
	}
	fmt.Println("a: ",a)
	
	// declaring number b
	b, err := bigint.NewInt("100")
	if err != nil {
		panic(err)
	}
	fmt.Println("b: ",b)

	// changing value of a with Set
	err = a.Set("200")
	if err != nil {
		panic(err)
	}
	fmt.Println("a: ",a)

	// adding a and b
	c := bigint.Add(a, b)
	fmt.Println("a+b = ",c)

	// subtracting a and b
	d := bigint.Sub(a, b)
	fmt.Println("a-b = ",d)
	
	// multiplying a and b
	e := bigint.Multiply(a, b) 
	fmt.Println("a*b = ",e)
	
	// find mod of a and b
	f := bigint.Mod(a, b)
	fmt.Println("Mod of a,b = ",f)

	// finding absolute value of given number
	abs, err := bigint.NewInt("-7")
	if err != nil {
		panic(err)
	}
	re := abs.Abs()
	fmt.Println("Abs val: ",re)
	
}
