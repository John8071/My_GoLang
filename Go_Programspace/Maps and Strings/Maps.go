/* Map is a built-in type in Go, to store key-value pairs
syntax:
make(map[type of key]type of value)
*/

package main

import "fmt"

type empl struct { // For Map of structs, empl struct definition
	salary  int
	country string
}

func main() {
	employeesalary := make(map[string]int)
	employeesalary["Stark"] = 1000000
	employeesalary["Rogers"] = 750000
	employeesalary["Banner"] = 900000

	fmt.Println("Employee salary list:", employeesalary)

	// can create a map during initialization itself

	emp_sal := map[string]int{
		"Hawkeye": 500000,
		"natasha": 750000,
	}
	fmt.Println("Emp_sal is:", emp_sal)

	//length of key

	fmt.Println("length of emp_sal is:", len(emp_sal))

	/* Zero value of a map is nil. If we try to add elements to such a map, run time panic error will occur, so we should intialise map before ading elements

	var e_sal map[string]int
	e_sal["steve"] = 12000
	fmt.Println("ësal is:", e_sal) //throw: panic: assignment to entry in nil map
	*/

	//Retrieving value from key

	employee := "Stark"
	salary := employeesalary[employee]
	fmt.Println("salary of", employee, "is", salary)

	// If there is no element, the map will return the zero value of the type of that element

	fmt.Println("Salary of joe is", employeesalary["joe"]) // output will be zero, as there is no element as joe

	// key to check element in map:    value, ok := map[key]

	newEmp := "josh"
	value, ok := employeesalary[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
		return
	}
	fmt.Println("element", newEmp, "not found!!")

	//iterate over a map

	for key, value := range employeesalary {
		fmt.Printf("Employee salary of %s = %d\n", key, value)
	}

	// Delete elements from key
	// syntax: delete(map, key)

	fmt.Println("Before Delete:", employeesalary)
	delete(employeesalary, "Banner")
	fmt.Println("\nAfter Delete:", employeesalary)

	// If we try to delete a key that is not present in the map, there will be no runtime error.
	delete(employeesalary, "Jigen")
	fmt.Println("After Delete false element:", employeesalary) // No error will be thrown, prints all elements in the map

	// Map of structs, struct employee is defined above main func

	emp1 := empl{
		salary:  12000,
		country: "India",
	}

	emp2 := empl{
		salary:  13000,
		country: "USA",
	}

	emp3 := empl{
		salary:  14000,
		country: "Japan",
	}

	employeeInfo := map[string]empl{
		"Arun":   emp1,
		"Mark":   emp2,
		"Itachi": emp3,
	}

	for name, info := range employeeInfo {
		fmt.Printf("\nEmployee name: %s, salary: %d, Country: %s", name, info.salary, info.country)
	}

	// Maps are called by reference

	fmt.Println("\nBefore Modification:", employeesalary)
	modified := employeesalary
	modified["Stark"] = 15000000
	fmt.Println("After modification:", employeesalary)

	// Maps cannot be compared with == operator, we have to compare them element by element
}
