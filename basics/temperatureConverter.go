package main
import "fmt"

func main(){
	var celcius float64 = 37.0

	var fahrenheit float64;

	fahrenheit = (celcius * 9/5) + 32;

	fmt.Printf("value %v , Type %T", celcius, celcius);
	fmt.Printf("value %v , Type %T", fahrenheit, fahrenheit);
}