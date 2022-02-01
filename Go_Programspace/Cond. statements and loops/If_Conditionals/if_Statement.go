// If Condition

package main

import "fmt"

func main() {

	var a int

	fmt.Println("Enter a number:")

	fmt.Scan(&a)

	if a > 10 {
		fmt.Println("a is greater!!")
		return // Ends the execution of the func. (here main()) and return to the caller
	}

	fmt.Println("a is lesser!!")
}
