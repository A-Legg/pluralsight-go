package main

import (
	"fmt"
	"os"
)

var (
	name, course string
	module       float64
)

func main() {

	name := os.Getenv("USER")
	course := "Go Fundamentals"

	fmt.Println("\nHi", name, "you're currently wantching",
		course)

	changeCourse(&course)
	fmt.Println("\nYou are now watching course", course)
}

func changeCourse(course *string) string {
	*course = "First Look: Native Docker Clustering"

	fmt.Println("Trying to change your course to", *course)

	return *course
}
