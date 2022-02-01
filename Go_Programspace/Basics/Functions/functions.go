// Functions in Go

/*
func functionname(parametername1 type, parametername2 type,...) returntype {
 //function body
}
*/

package main

import "fmt"

func total(goods int, cost float64) float64 { // also total(price, cost int), since both param. as same type
	res := float64(goods) * cost
	return res
}

func cal(length, breadth int) (int, int) { // return two values
	area := length * breadth
	perimeter := 2 * (length + breadth)
	return area, perimeter
}

/*func cal(l, b int) (a, p int) { // named return values
	area := l * b
	perimeter := 2 * (l + b)
	return
}*/

func main() {
	goods := 47
	price := 12.50
	tot := total(goods, price)

	fmt.Println("total amt is:", tot)

	l, b := 50, 25
	ar, peri := cal(l, b) // get two values from func.
	a, _ := cal(10, 20)   // _ is called blank identifier, here it will discard the second return value i.e 20
	// cal() gives two return values but you need only one so here we are using blank identifier where we need

	fmt.Println("\narea and perimeter is:", ar, peri)
	fmt.Println("\nArea is:", a)

}
