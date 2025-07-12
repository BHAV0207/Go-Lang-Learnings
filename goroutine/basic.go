package main 

import "fmt"
import "time"

func printMsg(){
	fmt.Println("hello from go routine ")
}


func main(){
	go printMsg();
	fmt.Println("hello worrld");
	time.Sleep(1 * time.Second) // wait for goroutine to finish
}

//  go routine add concurrency in the code , that is both the codes currently in the main function are executed concurrently that is together at the same time 
//  we add time because main function doesnt wait for the go routine to execute thus we add time so that go routine finished executing 
//  if we dont use the keyword go then it will become a sequencial call that is first the first function will execute then the second function will execute 