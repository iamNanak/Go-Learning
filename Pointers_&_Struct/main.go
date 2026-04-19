package main

import "fmt"

// Structs  -->> custom data type , we can modify it using pointers
type User struct {
	ID     int
	Name   string
	Salary float32
}

// receiver type
func (u *User) changeSalary(sal float32) {
	u.Salary = sal
}

type Rectangle struct {
	width  int
	height int
}

func changeValue(x int) {
	x = 20
}

func changeValueWithPointer(x *int) {
	*x = 20
}

func main() {
	x := 10

	changeValue(x)
	fmt.Println("value of x after changeValue is: ", x)

	changeValueWithPointer(&x)
	fmt.Println("value of x after changeValueWithPointer is: ", x)

	p := &x

	fmt.Println("value of x is: ", x)
	fmt.Println("value of p (address of x) is: ", p)
	fmt.Println("value of *p (dereferenced value of p) is: ", *p)

	*p = 60

	fmt.Println("new value of x is: ", x)

	// value of x is:  10
	// value of p (address of x) is:  0xc000012108
	// value of *p (dereferenced value of p) is:  10
	// new value of x is:  60
	//
	// If we need to update at any particular address then, we use pointers

	p1 := User{2, "Nanak", 2000000.00}

	p1.changeSalary(30000000)

	fmt.Println(p1)
	p2 := &p1

	p2.ID = 1
	fmt.Println(p1)

	r := Rectangle{10, 15}

	fmt.Println("Areas is: ", r.Area())

	r.Scale(2)
	fmt.Println("Area after Scale increase: ", r.Area())

}

// Value Receiver --> passes the copy of struct not original
func (r Rectangle) Area() int {
	return r.width * r.height
}

// Pointer Receiver --> passes original struct and original struct get change
func (r *Rectangle) Scale(factor int) {
	r.width = r.width * factor
	r.height *= factor
}
