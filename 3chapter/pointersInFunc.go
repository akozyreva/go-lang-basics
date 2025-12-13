package main

import "fmt"

/*
	a     = 10
	&a    = 0x1234
	p     = 0x1234
	*p    = 10
*/

func findFlat(isFound *bool) { // pass here pointer to var
	*isFound = true // change it's value
}

// Box it's like obj in JS
type Box struct {
	Items int
}

func newBox() *Box {
	b := Box{Items: 10}
	return &b
}

func main() {
	isFound := false
	findFlat(&isFound)   // send addr of var isFound
	fmt.Println(isFound) // isFound is changed!

	box := newBox()
	box.Items = 20
	fmt.Println(box.Items) // 20
}
