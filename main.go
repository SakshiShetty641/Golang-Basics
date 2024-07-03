package main

import (
	"fmt"
	"myapp/doctor"
)

func main() {

	// Variables -> Variables can be declared using the var keyword or the shorthand :=
	var a int = 10
	age := 20

	fmt.Println("Hello World!")
	sayHelloWorld("Hello Programming")

	var toSay = doctor.Intro()
	fmt.Println(toSay)
	fmt.Println(a, age)
}

// Introducing new functions
func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
