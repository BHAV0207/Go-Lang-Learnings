package main;

import "fmt";

func main() {
	var score int = 70;

	if score < 40 {
		fmt.Println("Failed");
	}else if score >= 40 && score < 75 {
		fmt.Println("Passed");
	}else{
		fmt.Println("Distinction");
	}
}