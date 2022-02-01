//Pointers are variables which stores the address of the another variable

package main

import "fmt"

func change_ptr(m *int) {
	*m = 65
}

func ptr_ret() *int {
	i := 25
	return &i
}

func main() {
	var b int = 2
	var a *int = &b //declaring a pointer 'a' and assigning address of b. '&' operator used to get address
	fmt.Printf("Type of a: %T\n", a)
	fmt.Println("Address:", a, "\nValue:", b)

	// zero value of a pointer is nil

	// create pointers using new()

	npointer := new(int)
	fmt.Printf("npointer type:%T, value:%d, address:%v\n", npointer, *npointer, npointer)
	*npointer = 100
	fmt.Printf("npointer type:%T, value:%d, address:%v\n", npointer, *npointer, npointer)

	c := &b //Short Hand Declaration of pointer
	fmt.Printf("type of c:%T\n", c)
	fmt.Println("Address:", c)

	fmt.Println("Value of b before:", b)
	*a++
	fmt.Println("Value of b after:", *a)

	// Passing pointer to a func.

	var n int = 45
	fmt.Println("Value of n before func. call:", n)
	n1 := &n
	change_ptr(n1)
	fmt.Println("value after func call:", n)

	// Returning pointer from func. call
	s := ptr_ret()
	fmt.Println("Value of s:", *s)

	// Do not pass a pointer to an array as an argument to a func. Use slice instead
	// Go does not support pointer arithmetic
}
