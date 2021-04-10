package main

import "fmt"

type person struct {
	age  int
	name string
}

// modifyFails tries to modify values of arguments passed to it.
func modifyFails(i int, s string, p person) {
	fmt.Println("in function before mutation:", i, s, p.name, p.age)
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
	p.age = 25
	fmt.Println("in function:", i, s, p.name, p.age)
}

// modMap modifies a map parameter.
func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

// modSlice modifies a slice parameter.
func modSlice(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
	s = append(s, 10)
	fmt.Println("s in function:", s)

}

func main() {
	p := person{}
	i := 2
	s := "Hello"
	modifyFails(i, s, p)
	fmt.Println(i, s, p)

	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m)

	intSlice := []int{1, 2, 3}
	modSlice(intSlice)
	fmt.Println(intSlice)
}
