package main;

import "fmt";

func main() {
	mapping  := map[string]int{
		"apple" : 10,
		"banana": 20,
		"cherry": 30,
	}

	fmt.Println("Original map:", mapping)

	mapping["banana"] = 25 // Update value for key "banana"
	delete(mapping, "cherry") // Remove key "cherry"

	for key , value := range mapping {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}