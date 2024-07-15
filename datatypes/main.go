package main

import "fmt"

func main() {

	// Basic types
	var isActive bool = true
	var age int = 30
	var pi float64 = 3.14159
	var message string = "Hello, Go!"
	var b byte = 'a'
	var r rune = '中'

	// Composite types
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	var scores []int = []int{98, 92, 85}
	var person map[string]string = map[string]string{"name": "Sakshi", "city": "Bangalore"}

	type Person struct {
		Name string
		Age  int
	}

	var p Person = Person{Name: "Tom", Age: 25}

	// Reference types
	var x int = 42
	var pPointer *int = &x

	ch := make(chan int)

	fmt.Println(isActive, age, pi, message, b, r)
	fmt.Println(numbers, scores, person)
	fmt.Println(p)
	fmt.Println(x, pPointer)
	fmt.Println(ch)

	a, b := 10, 10

	//** ARITHMETIC OPERATORS **//
	//  +       sum
	// -        difference
	// *        product
	// /        quotient
	// %        remainder
	// there is no power operator in Go. Use math.Pow(a, b) for raising to a power.

	fmt.Println(a + 5)
	fmt.Println(10 - b)
	fmt.Println(a * a)
	fmt.Println(a / a)
	fmt.Println(11 / 5)

	// Go is a Strong Typed Language
	// fmt.Println(a * b)       // =>  invalid operation: a * b (mismatched types int and float64)
	fmt.Println(a * int(b)) // => 50

	// IncDec Statements
	// The "++" and "--" statements increment or decrement their operands by the untyped constant 1.
	c := 10
	c++ // x is 11. Same as: x += 1
	c-- // x is 10. Same as: x -= 1

	//** ASSIGNMENT OPERATORS **//
	//  =   (simple assignment)
	// +=   (increment assignment)
	// -=   (decrement assignment)
	// *=   (multiplication assignment)
	// /=   (division assignment)
	// %=   (modulus assignment)

	a = 10
	a += 2 // => a is 12
	a -= 3 // => a is 9
	a *= 2 // => a is 18
	a /= 3 // => a is 6
	a %= 5 // => a is 1

	//** COMPARISON OPERATORS **//
	//  ==      equal values
	// !=       not equal
	// >        left operand is greater than right operand
	// <        left operand is less than right operand
	// >=       left operand is greater than or equal to right operand
	// <=       left operand is less than or equal to right operand

	fmt.Println(5 == 6)
	fmt.Println(5 != 6)
	fmt.Println(10 > 10)
	fmt.Println(10 >= 10)
	fmt.Println(5 < 5)
	fmt.Println(5 <= 5)

	//** LOGICAL OPERATORS **//
	// &&       logical and
	// ||       logical or
	// !        logical negation

	fmt.Println(0 < 2 && 4 > 1)
	fmt.Println(1 > 5 || 4 > 5)
	fmt.Println(!(1 > 2))

}
