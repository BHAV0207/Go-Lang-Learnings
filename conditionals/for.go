package main 
import "fmt";

func main(){
	
	for i := 1 ; i<=10 ; i++{
		fmt.Println(i)
	}

	var val int = 5
	for val > 0{
		fmt.Println(val)
		val--
	}

	str := []string{"Go", "is", "fun"}
	for _, value := range str{
		fmt.Println(value)
	}
}	