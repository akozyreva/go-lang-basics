package main

import "fmt"

type Switch string

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

type Toggable interface {
	toggle()
}

func main() {
	s := Switch("on")
	fmt.Println(s)
	s.toggle() // value was changed, because we send copy to method
	fmt.Println(s)
	var t Toggable = &s // you need to assign addr of s, because there are implemented with pointers methods
	t.toggle()
}
