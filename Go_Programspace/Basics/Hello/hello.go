package main

//The package “main” tells the Go compiler that the package should compile as an executable program
//instead of a shared library.
//Every executable Go application must contain the main function.
//This function is the entry point for execution. The main function should reside in the main package.

import "fmt" // Importing formatting package fmt

func main() {

	//main func. from where prog. execution starts, this is the case only when package main is used,
	// otherwise it is treated as a normal func. named main

	fmt.Println("Hello World!") //using Println() from fmt package.
	//  It returns the number of bytes written and any write error encountered.
}
