// Constants
//***** string constants *****

package main

import (
	"fmt"
)

func main() {

	const a = 10 // declaring a variable as a constant

	fmt.Println("a is:", a)

	fmt.Printf("a type is: %T\n", a)

	//a = 25 // Error, we cannot reassign values to a const var, throw: cannot assign to a (declared const)

	const ( // declaring multiple var. with diff. data types as constant
		name   = "John"
		age    = 26
		salary = 25000.00
	)

	fmt.Println("Name =", name, "\nAge =", age, "\nsalary =", salary)

	// Value of the const should be known at compile time

	//var x = math.Sqrt(4) // allowed since a is a var
	//const y = math.Sqrt(4) // not allowed, y is const and the value must be known at compile time, math.Sqrt() evaluated during runtime
	// throw: const initializer math.Sqrt(4) is not a constant

	//***** string untyped and types constants *****

	const n = "good morning" // untyped constant, have datatype associated with them but supply it only when code demands it

	var s = n // n default type is str, so s takes the same

	fmt.Printf("\ntype of n: %T\n", n)

	fmt.Printf("type of s: %T\n", s)

	fmt.Println("s is:", s)

	const h string = "hello" // typed constant

	/*
		var defaultName = "Sam"         //allowed, default type is string
		type myString string            // creating a new type called myString, which is of string type
		var customName myString = "Sam" //allowed, since sam is untyped const
		customName = defaultName        //not allowed, for Go compiler, string and myString are not same
	*/

	// ***** Boolean constants *****
	// same rules as of string constant (typed & untyped)

	const tconst = true // untyped

	type myBool bool

	var dconst = tconst

	fmt.Println("\ndconst: ", dconst)

	var nconst myBool = tconst // typed

	fmt.Println("\nnconst: ", nconst)

	//nconst = dconst // not allowed

	// ***** Numeric constants *****

	const c = 5 // untyped

	var intVar int = c // c acts as int

	var int32Var int32 = c // c acts as int32

	var float64Var float64 = c // c acts as float, 5.00

	var complex64Var complex64 = c // c acts as complex number without imaginary part, (5 + 0i)

	fmt.Println("\nintVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	// Numeric constants are allowed to be mixed into expressions

	var exp = 25.50 / 5 // 25.05 is float & 5 is int, since both are numeric const, this is allowed

	fmt.Printf("\ntype of exp:%T", exp)

	fmt.Println("\nexp:", exp)
}
