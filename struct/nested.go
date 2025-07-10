package main

import "fmt"

type Category struct {
	Name   string
	Parent string
}

type Product struct {
	Name     string
	Price    float64
	InStock  bool
	Category Category
}

func main() {
	var prod1 Product = Product{
		Name:    "laptop",
		Price:   100000.00,
		InStock: true,
		Category: Category{
			Name:   "Electronic",
			Parent: "someone",
		},
	}

	fmt.Println(prod1)
}
