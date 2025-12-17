package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	// time.Now()
	// time - package name
	// Now - method, which can be exported. it returns time.Time object??
	// time.Time -> Time is the type of package time
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(reflect.TypeOf(now))
	fmt.Println(year)

	// Add your code below:
	fmt.Print("Print", "is", "different")
	fmt.Print("See?")
	//PrintisdifferentSee? no spaces, no new row
}
