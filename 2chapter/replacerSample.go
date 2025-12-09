package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks!"
	/*
			strings.replaceAll exists as well, but
			NewReplacer is faster
			can be written many pairs, like:
		replacer := strings.NewReplacer("#", "e", ",", "!")
		•	"#" → "o"
		•	"!" → "?"
	*/
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
