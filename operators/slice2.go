package main 

import "sort"
import "fmt"

func main() {
	var sl []int = []int{1 , 2, 34 ,5 ,6 ,7}
	var sub []int = sl[2:5];

	fmt.Println(len(sub))
	fmt.Println(cap(sub))


	sort.Ints(sl)
	fmt.Println(sl)
}