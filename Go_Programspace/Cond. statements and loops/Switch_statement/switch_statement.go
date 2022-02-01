// switch statement

/*
A switch is a conditional statement that evaluates an expression and
compares it against a list of possible matches and executes the corresponding block of code
*/

package main

import "fmt"

func main() {

	var day string

	fmt.Println("Enter a working day:")
	fmt.Scan(&day)

	switch day {
	case "Monday":
		fmt.Println("Today is Monday!!")
	case "Tuesday":
		fmt.Println("Today is Tuesday!!")
	case "Wednesday":
		fmt.Println("Today is Wednesday!!")
	case "Thursday":
		fmt.Println("Today is Thursday!!")
	//case "Thursday":
	//	fmt.Println("Today is Thursday!!")	// Duplicate cases not allowed, throws error
	case "Friday":
		fmt.Println("Today is Friday!!")
	default: // Prints the below output for any other input other than given in switch cases, default case can be placed anywhere, not necessarily last
		fmt.Println("Please enter a working day!!!")
	}

	// Multiple expression in switch cases

	letter := "i"
	fmt.Printf("Letter %s is a ", letter)
	switch letter {
	case "a", "e", "i", "o", "u": //multiple expressions in case
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}

	// Expressionless switch cases

	num := 75
	switch { // expression is omitted
	case num >= 0 && num <= 50:
		fmt.Printf("%d is greater than 0 and less than 50", num)
	case num >= 51 && num <= 100:
		fmt.Printf("%d is greater than 51 and less than 100", num)
	case num >= 101:
		fmt.Printf("%d is greater than 100", num)
	}

	// Fallthrough in switch cases

	// In switch case, once the case is executed, execution will come out of switch.
	// If Fallthrough is used at the end of the case, the case next to the current case will also get executed
	// It will work even if next case is false
	// It cannot be used at the end of last case, otherwise compiler will throw error "cannot fallthrough final case in switch"

	var s int = 2

	switch s {
	case 1:
		fmt.Println("qwerty")
	case 2:
		fmt.Println("asdfgh")
		fallthrough // Fallthrough must be given at the end of the case, otherwise compiler will complain that it is out of place
	case 3:
		fmt.Println("zxcvbn")
	}

	// We can use break keyword to break during switch case execution, as we are doing in loop statements
}
