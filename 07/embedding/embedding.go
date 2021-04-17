package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

// Description returns a string representation of an Employee.
func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

// FindNewEmployees does stuff.
func (m Manager) FindNewEmployees() []Employee {
	// some biz logic here
	return []Employee{}
}

type Inner struct {
	X int
}

type Outer struct {
	Inner
	X int
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []Employee{},
	}

	fmt.Println(m.ID) // Embedding the Employee type in Manager type
	fmt.Println(m.Description())

	o := Outer{
		Inner: Inner{
			X: 10,
		},
		X: 20,
	}
	fmt.Println(o.X)
	fmt.Println(o.Inner.X) // Have to access Inner explicitly because field
	// has the same name
}
