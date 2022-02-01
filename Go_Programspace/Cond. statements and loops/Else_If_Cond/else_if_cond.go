// else if condition

package main

import "fmt"

func main() {

	var a int

	a = 67

	if a > 90 {
		fmt.Println("a is greater than 90!")
	} else if a > 10 && a <= 90 {
		fmt.Println("a is between 10 and 91")
	} else {
		fmt.Println("a is lessthan 10!!!")
	}
}
