// Slices are powerful and flexible wrappers around arrays, they dont own data, but represent data underlying

package main

import "fmt"

func slicefunc(sf []int) {
	sf[0] = 2
}

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	fmt.Println("a is:", a)
	var b []int = a[1:4] // slicing a part of array a, a[start:end-1]
	fmt.Print("b is:", b)

	c := [...]int{3, 4, 5}
	fmt.Println("\nc is:", c) //another way for creating slice

	my_arr := [...]int{11, 22, 33, 44, 55, 66, 77}
	my_slice := my_arr[1:4]
	fmt.Println("slice:", my_slice, "array:", my_arr)
	my_slice[0] = 23
	my_slice[1] = 34
	my_slice[2] = 45
	fmt.Println("slice:", my_slice, "array:", my_arr) // changes affect original array, as slice represent that array

	// When a number of slices share the same underlying array, the changes that each one makes will be reflected in the array

	my_slice2 := my_arr[1:4]
	my_slice2[2] = 46
	fmt.Println("slice:", my_slice2, "array:", my_arr)

	/*
		Length and Capacity of the slice

		length := number of elements in the slice
		capacity := number of elements in the underlying array starting from the index from which the slice is created
	*/

	fruits := [7]string{"apple", "orange", "mango", "Guava", "watermelon", "Grapes", "Avocado"}
	fruit_slice := fruits[0:4]
	fruit_slice2 := fruits[2:5]

	fmt.Printf("Length of slice: %d, Capacity of slice: %d", len(fruit_slice), cap(fruit_slice))
	fmt.Printf("\nLength of slice: %d, Capacity of slice: %d", len(fruit_slice2), cap(fruit_slice2))

	// Creating a slice using make func

	// syntax: func make( []type, length, capacity)

	// Default values of a slice will be zero, if it is created by using make func.

	i := make([]int, 5, 5)

	fmt.Println("\nslice by make:", i)

	// Append in the slice

	// When we append a slice, which is a representation of an array, in which we cannot extend, a new array is created and the values are copied

	s := []string{"suzuki", "mitsubishi", "subaru", "Ford"}
	s = append(s, "Toyota")
	fmt.Println("s is :", s) // output with a new element will be appended

	/*
		The zero value of a slice type is nil. A nil slice has length and capacity 0. It is possible to append values to a nil slice using the append function.
	*/
	var names []string //zero value of a slice is nil
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "tony stark", "bruce wayne", "steve rogers")
		fmt.Println("names contents:", names)
	}
	//Appending one slice to another

	veggies := []string{"potato", "tomato", "cauliflower", "onion"}

	food := append(fruit_slice, veggies...)

	fmt.Println("food:", food) // two slices appended to one

	slicefunc(my_arr[:])

	fmt.Println("AFter func:", my_arr)

}
