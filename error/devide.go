package main 

import (
	"fmt";
	"errors"

)
	



func devide (a , b float64)(float64 , error){
	if b == 0 {
		return 0 , errors.New("cannot devide by zero")
	}

	return a/b , nil
}


func main(){
	result , err := devide(10 ,3);

	if err != nil {
		fmt.Println("Error" , err)
		return
	}

	fmt.Println(result)
}