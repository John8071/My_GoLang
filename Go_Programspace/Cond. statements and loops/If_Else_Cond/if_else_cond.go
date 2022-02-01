// If else conditionals

/*
The else statement should start in the same line after the closing curly brace } of the if statement.
If not the compiler will complain.
The reason is because of the way Go inserts semicolons automatically.
In the rules, it's specified that a semicolon will be inserted after closing brace }, if that is the final token of the line.
So a semicolon is automatically inserted after the if statement's closing braces }


Since if{...} else {...} is one single statement, a semicolon should not be present in the middle of it.
Hence the program will fail to compile.
Therefore it is a syntactical requirement to place the else in the same line after the if statement's closing brace }.
*/

package main

import "fmt"

func main() {
	a := 12
	if a == 12 {
		fmt.Println("\nTRUE!!")
	} else {
		fmt.Println("\n FALSE!!")
	}

}
