package main

import "fmt"

/*
-> In Go, constants are immutable values that are known at compile time and do not change during the execution of the program.
-> They are declared using the const keyword.
-> Constants can be of various types, including numeric, string, and boolean.
-> Constants cannot be declared using the := syntax.
-> Constants must be initialized at the time of declaration.
-> Constants cannot be modified after declaration.
*/

func main() {
	/*mandatory to assign a value to a constant
		Error  when using constants is detected early at compile time
	    In group of constants if value is not assigned separately it picks from the previous one
	*/

	const days int = 7
	fmt.Println(days)

	const (
		min1 = 600
		min2
		min3
	)
	fmt.Println(min1, min2, min3)

}
