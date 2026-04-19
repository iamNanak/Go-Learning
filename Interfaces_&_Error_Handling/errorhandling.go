package main

import (
	"errors")



// func main(){
// 	result, err := divide(20, 0)
// 	if err != nil{
// 		fmt.Println("Error occured: ", err)
// 		return
// 	}

// 	fmt.Println("Result is --- ", result)
// }

func divide(a, b int)(int, error){
	if(b == 0){
		return 0, errors.New("Cannot divide by zero")
	}

	return a/b, nil
}
