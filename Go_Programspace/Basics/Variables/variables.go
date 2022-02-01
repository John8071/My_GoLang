package main

import ( // Importing Multiple packages
	"fmt"
	"math"
)

func main() {

	var age int // variable declaration, syntax: "var name type"

	//var age int = 10 // assign value during declaration

	//var age = 10 // since variable has been assigned an value initially, compiler infer the type from the value

	//fmt.Println("age is:", age) // output will be zero, when var is unintialised

	age = 10 // assign value to variable

	fmt.Println("age is:", age)

	//***** Multivar declaration *****

	//var height, weight int

	var height, weight = 100, 50 // Assign values to multiple variables

	fmt.Println("Height is:", height, "Weight is:", weight)

	//***** Multivar with different data types *****

	var (
		my_name   = "john"
		my_age    = 26
		my_height = 180
	)

	fmt.Println("Name = ", my_name, "Age = ", my_age, "Height = ", my_height)

	var x, y = 123.45, 178.90

	var z = math.Min(x, y) // using a func from a package, value is calculated in runtime and gets stored in z

	fmt.Println("Min value between x and y is ", z)

}
