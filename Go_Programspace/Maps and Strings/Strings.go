/*
Strings are different implementation in Go, compared to other languages

A string is a slice of bytes in Go.
Strings can be created by enclosing a set of characters inside double quotes " ".
Strings in Go are Unicode compliant and are UTF-8 Encoded.
*/

package main

import (
	"fmt"
	"unicode/utf8"
)

func printChar(s string) {
	fmt.Println("\nIn Char:")
	for i := 0; i < len(s); i++ { // Print char by char
		fmt.Printf("%c ", s[i])
	}
}

func Print_withRunes(s string) {
	runes := []rune(s)
	fmt.Println("\nWith runes, In Char:")
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func iterate_rune(s string) {
	fmt.Println("\nIterating rune with range:")
	for index, rune := range s {
		fmt.Printf("%d th element %c\n", index, rune)
	}
}

func mutate(s []rune) string {
	s[0] = 'W'
	return string(s)
}

func main() {
	string1 := "Good Morning" // defining a string
	fmt.Println(string1)
	printChar(string1) // func call to print char by char

	name := "Señor"
	fmt.Println("\nname:", name)
	printChar(name) // output : S e Ã ± o r

	/*
		we are trying to print the characters of Señor and it outputs S e Ã ± o r which is wrong.
		The reason is that the Unicode code point of ñ is U+00F1 and its UTF-8 encoding occupies 2 bytes c3 and b1.
		We are trying to print characters assuming that each code point will be one byte long which is wrong.
		In UTF-8 encoding a code point can occupy more than 1 byte.

		To solve this we are going to use runes

		A rune is a builtin type in Go and it's the alias of int32.
		Rune represents a Unicode code point in Go.
		It doesn't matter how many bytes the code point occupies, it can be represented by a rune

	*/
	Print_withRunes(name)

	// iterate rune with range
	name2 := "Señorita"
	iterate_rune(name2) // since ñ takes 2 bytes i.e. here 2 and 3, next element takes index 4

	/* String length

	len() returns the number of bytes
	some Unicode characters have code points that occupy more than 1 byte.
	Using len to find out the length of those strings will return the incorrect string length.

	The RuneCountInString(s string) (n int) function of the utf8 package can be used to find the length of the string.
	This method takes a string as an argument and returns the number of runes in it.
	*/

	str1 := "perfect"
	fmt.Println("String is:", str1)
	fmt.Println("Bytes with len():", len(str1))
	fmt.Println("Length with rune count():", utf8.RuneCountInString(str1))

	str2 := "Señor"
	fmt.Println("String is:", str2)
	fmt.Println("Bytes with len():", len(str2))
	fmt.Println("Length with rune count():", utf8.RuneCountInString(str2))

	// create a string from a slice of runes

	runeslice := []rune{'q', 'w', 'e', 'r', 't', 'y'}
	str_r := string(runeslice)
	fmt.Println("String from runeslice:", str_r)

	// string comparison

	if str1 == str2 {
		fmt.Println("\nTrue")
	} else {
		fmt.Println("\nFalse")
	}

	// string concatenation

	my_str1 := "Good"
	my_str2 := "Evening"
	res := my_str1 + "" + my_str2 // Method 1
	fmt.Println("\nConcatenated res:", res)

	//The Sprintf function formats a string according to the input format specifier and returns the resulting string
	res1 := fmt.Sprintf("%s %s", my_str1, my_str2) // Method 2
	fmt.Println(res1)
	//This format specifier takes two strings as input and has a space in between.
	//This will concatenate the two strings with a space in the middle.
	//The resulting string is stored in result

	// Strings are immutable in Go
	// but we can make a rune slice out of the string and change any element

	my_str3 := "Tell"
	fmt.Println("\nbefore mutate:", my_str3)
	fmt.Println("After mutate:", mutate([]rune(my_str3)))

}
