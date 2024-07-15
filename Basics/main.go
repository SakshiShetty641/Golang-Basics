package main

import "fmt"

func main() {

	//1st way
	var num int = 30
	fmt.Println("Number is ", num)

	//2nd way
	var name = "Tom"
	fmt.Println("Name is ", name)

	//3rd way
	age := 28 // using short declaration:
	fmt.Println("Age  is ", age)

	details := "Undefined"
	_ = details //order to avoid compile time error saying not used

	//Multiple declarations
	car, cost := "BMW", 1000
	fmt.Println(car, cost)

	//OTHER WAY for multiple declarations
	var (
		name1  string
		gender bool
		age1   int
	)
	fmt.Println(name1, gender, age1)

	var i, j int
	i, j = 10, 20
	_, _ = i, j // used to mute the variables

	j, i = i, j // to swap the variables
	fmt.Println(i, j)

	//expressions in short hand declarations
	sum := 3 + 4
	fmt.Println(sum)

	//Zero Values
	var a, b, c int
	fmt.Println(a, b, c)

	// An uninitialized variable or empty variable  will get the so called ZERO VALUE
	// The zero-value mechanism of Go ensures that a variable always holds a well defined value of its type
	var value int                          // initialized with 0
	var price float64                      // initialized with 0.0
	var name2 string                       // initialized with empty string -> ""
	var done bool                          // initialized with false
	fmt.Println(value, price, name2, done) // -> 0 0.0 ""  false

}
