// go run readfile.go
// 71.8go run readfile.go
// 56.2
// 89.5
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// open file on read
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file) // this obj is required for reading
	for scanner.Scan() {              // error can be here, but we don't read it here
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil { // here we check it
		log.Fatal(err)
	}

}
