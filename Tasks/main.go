package main

import (
	"fmt"
)

type Person struct{
	name string 
	surname string
	age int
	city string
}
  
func (p Person) greet() string {
	return "hello my name is " + p.name + " and I'm from  " + p.city 
}

//because its not  returning anything   
func (p *Person) hasbirthday() {
	p.age ++
}

func main(){


	//structs
	person1:=Person{name: "John", surname: "Smith", age: 43, city: "boston"}
	
	person1.hasbirthday()
	
	fmt.Println(person1.age)

	fmt.Println(person1.greet())


	// ids := []int{23,23,342,43,32,432,2,3,1321}
	
	// fmt.Println("listing all items ...")
	
	// for _,id := range ids{
	// 	fmt.Printf("ID:%d\n",id)
	// }
	
	// sum:=0
	// for _, s:= range ids{
	// 	sum=s+sum
	// }
	
	// println(sum)
	


	// for i := 0; i < 100; i++ {
		
	// 	if i%15==0{
	// 		fmt.Println("fizzbuzz")
	// 	}else if i%5==0{
	// 		fmt.Print("fizz")
	// 	}else if i%3==0{
	// 		fmt.Println("buzz")
	// 	}else
	// 	{
	// 		fmt.Println(i)
	// 	}
	// }
} 

