// Type conversion

package main

import "fmt"

func main() {
	i := 5
	j := 67.98

	fmt.Printf("i is of type: %T\n", i)
	fmt.Printf("j is of type: %T\n", j)

	sum := i + int(j) // type converion from float to int for j

	fmt.Println("sum: ", sum)
	fmt.Printf("sum is: %T\n", sum)

	value := j + float64(i) // typr conversion from int to float for i

	fmt.Println("value: ", value)
	fmt.Printf("value is: %T\n", value)

	a := 10
	var b float64 = float64(a) //this statement will not work without explicit conversion
	fmt.Println("b is:", b)

}
