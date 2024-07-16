package main

import "fmt"

func main() {
	var a int = 4
	var b float64 = 0.98
	var c bool = false
	var d string = "Hello"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	x := 20
	y := 15.5
	z := "Gopher!"
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	var (
		f int
		g float64
		h bool
		i string
	)

	f, g, h, i = 8, 0.9, true, "Hello"
	fmt.Println(f, g, h, i)

	u, v, w := 4, 7, 8
	fmt.Println(u, v, w)

	r := 8
	_ = r // to mute r

	//Exercise -> find error
	var p float64 = 7.1

	q, s := true, 3.7

	p, q = 5.5, false

	_, _, _ = p, q, s

	//Exercise
	name := "Golang"
	fmt.Println(name)

	const daysWeek int = 7
	const lightSpeed float64 = 299792458
	const pi float64 = 3.14159
	fmt.Println(daysWeek, lightSpeed, pi)

	const (
		daysOfWeek  = 7
		lightsSpeed = 87666778
		valueOfPi   = 3.14159
	)
	fmt.Println(daysOfWeek, lightsSpeed, valueOfPi)

	const secPerDay int = 24 * 60 * 60
	const daysYear = 365
	fmt.Println(secPerDay * daysYear)

	//Exercise => error
	const t int = 10

	// declaring a constant of type slice int ([]int) => cannot be done
	var m = []int{1: 3, 4: 5, 6: 8}
	_ = m

	const e = 7
	const l float64 = 5.6
	const k = e * l

	//x := 8       cannot initiate a constant at runtime (constants belong to compile-time)
	//	const xc int = x

	//const noIPv6 = math.Pow(2, 128)  cannot initiate a constant at runtime (constants belong to compile-time)
}
