package main

import "fmt"

func main() {
	ages := make(map[string]string)

	ages["India"] = "delhi"
	ages["Australia"] = "Canbara"
	ages["spain"] = "idk"

	fmt.Println(ages["India"])
	val, exist := ages["Germany"]
	if exist {
		fmt.Println(val)
	} else {
		fmt.Println("capital of germany does not exist")
	}
}
