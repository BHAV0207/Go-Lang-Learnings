package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	people := []People{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 28},
	}

	jsonData, err := json.Marshal(people)
	if err != nil {
		fmt.Println("err")
		return
	}

	fmt.Println(string(jsonData))
}
