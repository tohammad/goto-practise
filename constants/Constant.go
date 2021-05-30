package main

import "fmt"

func main() {
	const myConstant int = 1 // If you make it caps , will get exported
	fmt.Printf("%v %T\n", myConstant, myConstant)
}
