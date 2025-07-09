package main 
import "fmt";

func areaParemeter(length int , width int) (area int , peremeter int){
	area = length*width;
	peremeter = 2*(length + width);

	return area , peremeter;
	//  we can also just write return and go will take care of the rest 
}


func main () {
	length := 10
	width := 5

	fmt.Println(areaParemeter(length , width))
}