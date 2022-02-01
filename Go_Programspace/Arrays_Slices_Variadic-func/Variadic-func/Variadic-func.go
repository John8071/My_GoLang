/*
A variadic function is a function that accepts a variable number of arguments.
If the last parameter of a function definition is prefixed by ellipsis ..., then the function can accept any number of arguments for that parameter.
Exmaple:
func hello(a int, b ...int){
}

hello (1, 2) // a gets 1 and b gets 2

hello (5, 6, 7, 8, 9) // b gets 6, 7, 8, 9 while a gets first one

hello(1) // a gets 1 but b is zero, this is also fine, its possible to pass zero to a variadic func.

Note:
Only the last parameter of a function can be variadic
Otherwise compiler will throw: syntax error: cannot use ... with non-final parameter
*/

package main

import "fmt"

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func find_new(s ...string) {
	s[0] = "Go"
	s = append(s, "Playground")
	fmt.Println(s)
}
func main() {
	w := []string{"hello", "World"}
	find_new(w...)
	fmt.Println("w is:", w)
	find(89, 90, 91, 89, 100)
	find(24, 24)
	find(5)
}
