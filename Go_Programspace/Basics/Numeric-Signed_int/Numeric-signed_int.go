// Numeric datatype, with signed integer

// Another Numeric datatype, rune is alias of int32

/*
int8: represents 8 bit signed integers
size: 8 bits
range: -128 to 127

int16: represents 16 bit signed integers
size: 16 bits
range: -32768 to 32767

int32: represents 32 bit signed integers
size: 32 bits
range: -2147483648 to 2147483647

int64: represents 64 bit signed integers
size: 64 bits
range: -9223372036854775808 to 9223372036854775807
*/

// In general case we will use simply int, unless otherwise required specifically for each type.

package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var a int = 25

	fmt.Println("a is", a)

	var x int32 = 89

	var y int64 = 95

	fmt.Println("value of a is", x, "and b is", y)

	fmt.Printf("type of x is %T, size of x is %d", x, unsafe.Sizeof(x)) //type and size of x, ==> int32, 4

	fmt.Printf("\ntype of y is %T, size of y is %d", y, unsafe.Sizeof(y)) //type and size of y, ==> int64, 8

}
