// Interfaces --> It like a contract of methods. If any type implements that methods, then it automatically satisfy that Interfaces

// type Name interface{}  -->> Empty interface

package main

import "fmt"


type Shape interface{
	Area() float64  // only signature of method
	Perimeter() float64
}

type Rectangle struct{
	width, height float64
}

type Circle struct{
	radius float64
}

func (r Rectangle) Area() float64{
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64{
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64{
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64{
	return 2 * 3.14 * c.radius
}

func main(){
	var s Shape
	s = Rectangle{10, 15}
	fmt.Println("Area of Rectangle is: ", s.Area())
	fmt.Println("Perimeter of Rectangle is: ", s.Perimeter())

	s = Circle{5}
	fmt.Println("Area of Circle is: ", s.Area())
	fmt.Println("Perimeter of Cicle is: ", s.Perimeter())

	result, err := divide(10, 0)
	if err != nil{
		fmt.Println("Error occured: ", err)
		return
	}

	fmt.Println("Result is --- ", result)

}
