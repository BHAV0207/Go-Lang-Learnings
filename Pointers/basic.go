package main 

import "fmt"


func main(){
	var x int = 10;
	var p *int = &x

	fmt.Println(x);
	fmt.Println(p);

	fmt.Println(*p)
	//  when you use * with a poniter it deferences the value and gives the output of the original value 

	x= 90;
	fmt.Println(x);


	*p = 99;
	fmt.Println(*p)
}