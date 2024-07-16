package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// if condition_that_evaluates_to_boolean{
	//      perform action1
	// }else if condition_that_evaluates_to_boolean{
	//      perform action2
	// }else{
	//      perform action3
	// }

	price, inStock := 100, true

	if price >= 80 { // parenthesis are no required to enclose the testing condition
		fmt.Println("Too Expensive")
	}

	if price <= 100 && inStock == true { //the same with: if price <= 100 && inStock { }
		fmt.Println("Buy it!")
	}

	// In Go there is not such a thing like the Truthiness of a variable.
	// Error:
	// if price {
	//  fmt.Println("We have price!")
	// }

	// only one if branch will be executed
	if price < 100 {
		fmt.Println("It's cheap!")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else { //executed only once if all the if branches are false (it's optional)
		fmt.Println("It's Expensive!")
	}

	//Getting the User input

	fmt.Println("os.Args:", os.Args) // os.Args is slice of strings ([]string)

	// accessing command line arguments using indexes
	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st Argument:", os.Args[1])
	fmt.Println("2nd Argument:", os.Args[2])
	fmt.Println("No. of items inside os.Args:", len(os.Args))

	//conversion to int
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Converted number:", num)
	}

	//Loops
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// has the same effect as a while loop in other languages
	// there is no while loop in Go
	j := 10
	for j >= 0 {
		fmt.Println(j)
		j--
	}

	// handling of multiple variables in a for loop
	for i, j := 0, 100; i < 10; i, j = i+1, j+1 {
		fmt.Printf("i = %v, j = %v\n", i, j)
	}

	// infinite loop
	// sum := 0
	// for {
	//  sum++
	// }
	// fmt.Println(sum) //this line is never reached

	//Break and Continue statement
	// The continue statement rejects all the remaining statements in the current iteration of the loop
	// and moves the control back to the top of the loop.
	// printing even numbers less than or equal to 10
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue // skipping the remaining code in this iteration
		}
		fmt.Println(i)
	}

	count := 0
	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Printf("%d is divisible by 13\n", i)
			count++
		}

		if count == 10 { //if 10 numbers were found, break!
			break //it breaks the current loop (inner loop if there are more loops)
		}
	}

	// the break statement is not terminating the program entirely;
	fmt.Println("Just a message after the for loop")

}
