// Get user input

/* %s -str, %d -int, %g -float, %t -bool */

package main

import "fmt"

func main() {

	/*
		The fmt.Scan() function in Go language scans the input texts which is given in the standard input,
		reads from there and stores the successive space-separated values into successive arguments.
		Moreover, this function is defined under the fmt package
	*/

	//var name string

	/*fmt.Println("Enter your name:")

	fmt.Scan(&name)

	fmt.Println("Name is:", name)*/

	/*
		The fmt.Scanf() function in Go language scans the input texts which is given in the standard input, reads from there and
		stores the successive space-separated values into successive arguments as determined by the format.
	*/

	/*var a int

	fmt.Println("\nEnter a number:")

	fmt.Scanf("%d", &a) // scanf() used with format specifiers

	fmt.Printf("Value of a is: %d", a)*/

	/*
		The fmt.Scanln() function in Go language scans the input texts which is given in the standard input,
		reads from there and stores the successive space-separated values into successive arguments.
		This function stops scanning at a newline and after the final item, there must be a newline or EOF.
		Moreover, this function is defined under the fmt package.
	*/

	var str_name string

	fmt.Println("Enter name:")

	fmt.Scanln(&str_name)

	fmt.Printf("Name: %s", str_name)

}
