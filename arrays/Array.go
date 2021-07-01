package main

import "fmt"

func main() {
	grades := [3]int{97, 85, 83}
	fmt.Printf("Grades: %v\n", grades)
	grades1 := [...]int{97, 85, 83}
	fmt.Printf("Grades 1: %v\n", grades1)
	var students [3]string
	students[0] = "Hammad"
	students[1] = "Mustafa"
	students[2] = "Lisa"
	fmt.Printf("Student: %v\n", students[2])
	fmt.Printf("Students Length: %v\n", len(students))
	fmt.Printf("Students Length: %v\n", cap(students))

	// slice
	grades3 := grades[2:]
	fmt.Printf("Grades: %v\n", grades3)

	// make
	grades4 := make([]int, 3)
	fmt.Printf("Grades: %v\n", grades4)

}
