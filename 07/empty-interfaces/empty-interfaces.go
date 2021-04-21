package main

import "fmt"

type LinkedList struct {
	Value interface{}
	Next  *LinkedList
}

// Insert adds a node to a LinkedList.
func (ll *LinkedList) Insert(pos int, val interface{}) *LinkedList {
	if ll == nil || pos == 0 {
		return &LinkedList{
			Value: val,
			Next:  ll,
		}
	}
	ll.Next = ll.Next.Insert(pos-1, val)
	return ll
}

// String stringifies a LinkedList
func (ll *LinkedList) String() string {
	if ll == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%v->%s", ll.Value, ll.Next)
}

func main() {
	var ll *LinkedList
	fmt.Printf("%v\n", ll)
	ll = ll.Insert(0, 10)
	fmt.Printf("%v\n", ll)
	ll = ll.Insert(0, "foo")
	fmt.Printf("%v\n", ll)
	ll = ll.Insert(1, 4.5)
	fmt.Printf("%v\n", ll)
	ll = ll.Insert(10, false)
	fmt.Printf("%v\n", ll)
	ll = ll.Insert(2, struct{}{})
	fmt.Printf("%v\n", ll)
}
