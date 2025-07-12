package main 

import "fmt";
import "time"

func printNum(n int){
	// time.Sleep(time.Duration(n) * 100 * time.Millisecond)
	fmt.Println(n)
}

func main (){
	for i:= 0 ;i<=5 ;i++{
		go printNum(i);
		time.Sleep(1*time.Second)
	}
}