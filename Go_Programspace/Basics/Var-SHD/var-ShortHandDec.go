// Short Hand Declaration of Variables
// Here no need to put var keyword before any variables, instead use := after var name

package main

import "fmt"

func main() {

	count := 10 // short hand declaration

	fmt.Println("count is :", count)

	name, age := "John", 26 // type is taken depending upon the value intialised

	fmt.Println("Name =", name, " Age =", age)

	// ==> Short hand dec. requires assignment for all var. declared on LHS

	//Height, Weight := 100 // Error

	//fmt.Println("Height =", Height, "Weight =", Weight) // throw: "assignment mismatch: 2 variables but 1 value"

	// ==> SHD requires atleast one new var. to be assigned, when using var repetitively

	// Example 1:

	a, b := 20, 30 // declare variables a and b

	fmt.Println("a is", a, "b is", b)

	b, c := 40, 50 // b is already declared but c is new

	fmt.Println("b is", b, "c is", c)

	b, c = 80, 90 // assign new values to already declared variables b and c

	fmt.Println("changed b is", b, "c is", c)

	//b, c := 100, 150 // Error, both are declared already, so it throw: no new variables on left side of :=

	//fmt.Println("b is", b, "c is", c)

	// Example 2:

	//a, b := 20, 30 //a and b declared

	//fmt.Println("a is", a, "b is", b)

	//a, b := 40, 50 //error, no new variables

	// ==> We cannot use same var for multiple data types

	num1 := 10

	// num1 := "john" // thow: cannot use "john" (untyped string constant) as int value in assignment

	fmt.Println("num1 =", num1)

}
