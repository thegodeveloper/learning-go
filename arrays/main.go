package main

import "fmt"

func main(){
	var a[5]int
	fmt.Println(a)
	fmt.Printf("a length: %d\n", len(a))

	b := [5]int{0,1,2,3,4}
	fmt.Println(b)
	fmt.Printf("b length: %d\n", len(b))

	c := [5]int{0,1,2}
	fmt.Println(c)
	fmt.Printf("c length: %d\n", len(c))
}
