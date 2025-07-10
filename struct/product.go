package main
import "fmt"


type Product struct {
	name string;
	price float64;
	inStock bool;
}


func  (p Product) display(){
	fmt.Printf("Product: %s\nPrice: %.2f\nIn Stock: %t\n", p.name, p.price, p.inStock)
}

func (p Product) stockByValue(status bool){
	p.inStock = status;
}

//  when you do not use pointer it is taken by value which makes a copy of the object 
// 	when you use pointer it is taken by reference whcih stirctly changes the object 
func (p *Product) stockByReference(status bool){
	p.inStock = status;
}


func main (){
	var product1 Product = Product{
		name: "Laptop",
		price: 36000.00,
		inStock: true,
	} 

	product1.stockByReference(false);
	product1.stockByValue(false);
	product1.display()
	fmt.Println(product1)
}



//  the difference between ---> func display() (p Product) is a stand alone function that returns a object of the type product 
//  while func (p Product) display()  is a type of reciever method which is defined for the type called product , like a function we creat inside a particular class only for the objects of that class in java 