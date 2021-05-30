package main

import (
	"fmt"
	"strconv"
)

// conversion

func main() {
	var i int = 32
	fmt.Printf("%v %T\n", i, i)
	var j float32
	j = float32(i)
	fmt.Printf("%v %T\n", j, j)

	var k string
	k = strconv.Itoa(i)
	fmt.Printf("%v %T\n", k, k)
}
