package greeting

import "fmt"

// Hello this method is for import, so it starts from capital letter
// only functions with capital letters are being exported
func Hello() {
	fmt.Println("Hello")
}

func Hi() {
	fmt.Println("Hi")
}

func AllGreetings() {
	// I can call them directly - without this or anythgin
	Hello()
	Hi()
}
