package main;
import "fmt"
import "strings"


func main() {
	text := "go is fun and go is powerful"
	words := strings.Split(text, " ")

	freq := make(map[string]int)

	for _ , word := range words {
		freq[word]++
	}

	fmt.Println(freq)
}