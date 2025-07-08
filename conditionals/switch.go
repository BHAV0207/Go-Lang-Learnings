package main 

import "fmt";

func main() {
//  in go lang there is no break statement in switch case
var day int = 13;

switch day {
case 1: 
	fmt.Println("Monday");
case 2:
	fmt.Println("Tuesday");
case 3:
	fmt.Println("Wednesday");
case 4:		

	fmt.Println("Thursday");
case 5:
	fmt.Println("Friday");
case 6 , 7:
	fmt.Println("Weekend");
default:
	fmt.Println("Invalid day");
}
}