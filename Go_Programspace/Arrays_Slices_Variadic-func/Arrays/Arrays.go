// Arrays in Go

//An array is a collection of elements that belong to the same type

package main

import "fmt"

func Change_ar(ar [4]int) {
	ar[0] = 3
	fmt.Println("Inside func:", ar)
}

func main() {

	var a [3]int // declaring an int array of length 3

	fmt.Println(a) // default value of an empty array is zero, so output is [0 0 0]

	a[0] = 1 // assigning values to each index in the array
	a[1] = 2
	a[2] = 3

	fmt.Println("The values are:", a)

	b := [3]int{10, 20, 30} // Short Hand Declaration of an array

	fmt.Println(b)

	fmt.Println("length of array:", len(b)) //len() func used to find the length of an array

	c := [3]int{40} // length is 3, but only one value is given, then remaining elements are filled with zero

	fmt.Println(c) // so the output is [40 0 0]

	d := [...]int{50, 60, 70} // [...] given so that compiler will determine the size of the array

	fmt.Println(d)

	// the size is an part of an array, hence it acts as a type for that array, so the below gives error

	/*
		a := [3]int{}
		var b [5]int
		b = a // gives error, throws: "cannot use a (type [3]int) as type [5]int in assignment"
	*/

	/*
		Array in Go is of value type, not reference type.
		If an array is assigned to a new variable, it is simply copied from original array.
		Even if we make changes, it wont affect the original array

		Same applies when passing it in a func., it is passed by value, not as reference
	*/

	m := [...]string{"India", "Germany", "USA", "Japan", "Russia"}
	n := m
	n[1] = "Singapore"

	fmt.Println("m is:", m) // remains same
	fmt.Println("n is:", n) // gives output with change

	num := [4]int{1, 2, 4, 5}
	fmt.Println("before func:", num)
	Change_ar(num)
	fmt.Println("After func:", num) // output doesn't change

	// iterate an array

	// we can use len() to iterate an array or we can use range keyword

	nu := [...]int{67, 89, 21, 78}
	for i, v := range nu { //range returns both the index and value
		fmt.Printf("%d th element of nu is %d\n", i, v)
		//fmt.Println("v is:", v)
	}
	fmt.Println("\nElements of nu:", nu)

	// multidimensional array

}
