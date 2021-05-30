package main

import "fmt"

var gi int = 190 // // lowercase first letter for package scope

func main() {
	var i int
	i = 7
	var j int = 8
	var k float32 = 8.9
	fmt.Println(i)
	fmt.Println(j)
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", k, k)
	fmt.Println(gi)
}
