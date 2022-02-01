//Numeric datatype, with complex numbers

/*
complex64: complex numbers which have float32 real and imaginary parts, float32 + float32 gives float64
complex128: complex numbers with float64 real and imaginary parts, float64 + float64 gives float128

syntax to declare a complex number:
1. short hand syntax, var := realpart + img.part i		Eg: c := 5 + 6i
2. function, func (r, i Float type) ComplexType		Eg: complex(5, 6)

Real and Img. parts must be of same type i.e either float32 or float64
*/

package main

import "fmt"

func main() {
	com1 := 5 + 6i
	com2 := complex(5, 6)
	com3 := com1 + com2
	fmt.Println("com3 is:", com3)

}
