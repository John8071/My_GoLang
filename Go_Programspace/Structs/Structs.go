// Struct - user-defined type which represents collection of fields

package main

import "fmt"

type student struct {
	Name         string
	Age          int
	Dept         string
	Average_Mark float64
}

type Person struct { //Anonymous fields
	string
	int
}

type Address struct {
	city  string
	state string
}

type person struct { //Nested Struct
	name string
	age  int
	Address
}

type name struct {
	firstName string
	lastName  string
}

func main() {
	stu1 := student{
		Name:         "Arun",
		Age:          19,
		Dept:         "EEE",
		Average_Mark: 87.72,
	}

	stu2 := student{"Baskar", 19, "ECE", 90.56} // defining struct without specifying filed names is also possible, but may confuse as which data belong to which field

	fmt.Println("Student 1 Details:", stu1)
	fmt.Println("Student 2 Details:", stu2)

	// Anonymous struct: declaring struct without creating new data type

	emp1 := struct {
		Name   string
		Age    int
		Salary int
	}{
		Name:   "James",
		Age:    26,
		Salary: 75000,
	}

	fmt.Println("\nEmployee details:", emp1)

	// Accessing Individual fields

	fmt.Println("Employee Name:", emp1.Name)
	fmt.Println("Employee Age:", emp1.Age)
	fmt.Println("Employee Salary:", emp1.Salary)

	emp1.Salary = 80000 // Altering data in a particular field
	fmt.Println("Employee Salary:", emp1.Salary)

	// Zero values of struct depends upon the individual field data type
	//It is also possible to specify values for some fields and ignore the rest. In this case, the ignored fields are assigned zero values.

	stu3 := student{
		Name: "Adam",
		Age:  22,
	}

	fmt.Println("Student 3 details:", stu3)
	fmt.Println("Student 3  Dept.:", stu3.Dept)               // Default is empty string
	fmt.Println("Student 3 Average Mark:", stu3.Average_Mark) // Default is zero

	//pointers to struct

	stu4 := &student{
		Name:         "Josh",
		Age:          22,
		Dept:         "CSE",
		Average_Mark: 92.76,
	}

	fmt.Println("Student 4 Name:", (*stu4).Name) // we can also give stu4.Name

	// Anonymous fields can be created in struct, in that case the field data type becomes the field name by default

	p1 := Person{
		string: "Joe",
		int:    34,
	}
	fmt.Println("Person 1 Details:", p1)

	//Nested Structs

	p2 := person{
		name: "Justin",
		age:  25,
		Address: Address{
			city:  "illinois",
			state: "chicago",
		},
	}

	fmt.Println("Person 2 name:", p2.name)
	fmt.Println("Person 2 age:", p2.age)
	fmt.Println("Person 2 city:", p2.Address.city)
	fmt.Println("Person 2 state:", p2.Address.state)

	// promoted fields
	//Here city and state are promoted fields, they belong to Address, but can be accessed as if they belong to struct person
	fmt.Println("Person 2 city:", p2.city)
	fmt.Println("Person 2 state:", p2.state)

	//Structs are value types and are comparable if each of their fields are comparable.
	//Two struct variables are considered equal if their corresponding fields are equal.
	//But, Struct variables are not comparable if they contain fields that are not comparable. It will throw error stating not comparable

	name1 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	name2 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	name4 := name{
		firstName: "Steve",
	}

	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}

}
