// Numeric datatype, with float

/*
float32: 32 bit floating point numbers
float64: 64 bit floating point numbers
*/

package main

import "fmt"

func main() {
	var a float32 = 58.08
	var b float64 = 98.67
	//a, b := 58.08, 98.67 // automatically take depending on system architecture in which it is running
	fmt.Printf("type of:\n a = %T\n b = %T\n", a, b)

}
