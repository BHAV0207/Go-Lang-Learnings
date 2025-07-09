package main

import "fmt";


func minMax(a int, b int) (int , int) {
	if a < b {
		return a , b
	}else{
		return b , a 
	}
}
func main() {
		var num1 int =  10;
		var num2 int = 5;

		var ans1 , ans2 int = minMax(num1 , num2);

		fmt.Println("min value is", ans1 , "max value is" , ans2)
}