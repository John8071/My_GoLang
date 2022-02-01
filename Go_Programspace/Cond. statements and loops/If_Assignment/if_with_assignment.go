// if condition with assignment

/*
Syntax:

if assignment-statement; condition {
}

assignment-statement is first executed before the condition is evaluated.
*/

package main

import "fmt"

func main() {

	var number int = 50

	if number := 10; number < 15 {

		fmt.Println("number is lesser than 15!!") // The value of number here will be 10 and its range is within if statment only

	}

	fmt.Println("Number is:", number)
}
