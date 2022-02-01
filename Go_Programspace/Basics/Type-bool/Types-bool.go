package main

import "fmt"

func main() {

	a, b := false, true

	c := a && b

	fmt.Println("c is", c) // In && operation, only true and true results in true

	d := a || b

	fmt.Println("d is", d) // In || operation, only false or false results in false
}
