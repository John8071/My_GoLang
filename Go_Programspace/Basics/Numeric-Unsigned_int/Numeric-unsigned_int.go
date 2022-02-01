// Numeric datatype, with unsigned integer

/*
uint8: represents 8 bit unsigned integers
size: 8 bits
range: 0 to 255

uint16: represents 16 bit unsigned integers
size: 16 bits
range: 0 to 65535

uint32: represents 32 bit unsigned integers
size: 32 bits
range: 0 to 4294967295

uint64: represents 64 bit unsigned integers
size: 64 bits
range: 0 to 18446744073709551615

uint : represents 32 or 64 bit unsigned integers depending on the underlying platform.
size : 32 bits in 32 bit systems and 64 bits in 64 bit systems.
range : 0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems
*/

//Another Numeric data type, byte is alias of uint8

package main

import "fmt"

func main() {
	var b uint = 87
	fmt.Println("b is", b)
}
