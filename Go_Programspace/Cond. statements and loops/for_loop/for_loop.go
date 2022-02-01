// for loop

/*
for is the only loop available in Go

syntax:

for initialisation; condition; post {
}

The initialisation statement will be executed only once
After the loop is initialised, the condition will be checked.
If the condition evaluates to true, the body of the loop inside the { } will be executed followed by the post statement.
The post statement will be executed after each successful iteration of the loop.
After the post statement is executed, the condition will be rechecked. If it's true, the loop will continue executing, else the for loop terminates.

*/

package main

import "fmt"

func main() {

outer: // label
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Println("i :", i, "j :", j)
			if i == j {
				break outer // break the loop from part mentioned by "outer" label
			} else {
				continue //break current iteration and move to next
			}
		}

	}

}
