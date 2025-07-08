package main

import "fmt"


func main(){
	var num int = 10;
	var str string = "hello";
	var boo bool = true;
	var fl float64 = 3.14;

	fmt.Printf("value %v , type %T" , num, num)
	fmt.Printf("value %v , type %T" , str, str)
	fmt.Printf("value %v , type %T" , boo, boo)
	fmt.Printf("value %v , type %T" , fl, fl)
}