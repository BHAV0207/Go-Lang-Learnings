package main
import "fmt"

func square(num int) int {
	return num *num
}


func main(){
	var num int = 4;

	fmt.Println(square(num));
}