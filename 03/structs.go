package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string // no commas! not a map!
	age  int
	pet  string
}

func main() {

	var fred person // all fields are zero values
	printPerson(fred)

	bob := person{}
	printPerson(bob)

	// empty struct literal and declaration without assignment are the same
	fmt.Println(reflect.DeepEqual(fred, bob)) // true

	julia := person{"Julia", 40, "cat"}
	printPerson(julia)

	// struct literal with named fields for clarity
	beth := person{age: 30, name: "Beth"}
	printPerson(beth)

	bob.name = "Bob"
	fmt.Println(beth.name)

	// Anonymous structs
	var person struct {
		name string
		age  int
		pet  string
	}
	person.name = "bob"
	person.age = 50
	person.pet = "dog"

	pet := struct {
		name string
		kind string
	}{name: "Fido", kind: "dog"}
	fmt.Println(pet)

	type firstPerson struct {
		name string
		age  int
	}
	// firstPerson type can be converted to secondPerson type because
	// field names, order, types, match
	type secondPerson struct {
		name string
		age  int
	}
	// cannot convert firstPerson to thirdPerson type
	// because field order does not match
	type thirdPerson struct {
		age  int
		name string
	}
	// cannot convert firstPerson to fourthPerson because field names don't
	// match
	type fourthPerson struct {
		firstName string
		age       int
	}
	// cannot convert firstPerson to fifthPerson because additional field
	type fifthPerson struct {
		name          string
		age           int
		favoriteColor string
	}

	// can compare structs WITHOUT TYPE CONVERSION if at least 1 is
	// anonymous and field names, order, and types match
	f := firstPerson{name: "Bob", age: 50}
	var g struct {
		name string
		age  int
	}
	g = f
	fmt.Println(f == g)
}

// printPerson ...
func printPerson(p person) {
	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p.pet)
}
