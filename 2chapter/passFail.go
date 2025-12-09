package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	/*
		ReadString returns output and err
		_ means - ignore err, I don't need it
	*/

	input, _ := reader.ReadString('\n') // enter text before client press enter
	// this code returns err, because 'err' var is not used!
	// input, err := reader.ReadString('\n')
	fmt.Println(input)
	fmt.Print("Enter a grade again: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		// msg about error and stop program
		// but in this case it's unreal, only possible to simulate
		log.Fatal(err)
	}
	input = strings.TrimSpace(input) // rm all trailing spaces
	fmt.Println(input)
	fmt.Println(reflect.TypeOf(input))
	// var status string it's also possible. then it will be empty
	status := "failing"
	if input >= "60" {
		a := 1
		status = "passing"
		// status := "new var" it won't work. it's new var inside of if block, which overlaps var outside
		fmt.Println(a)
	}
	// fmt.Println(a) it doesn't see a, because it's in block of if!!
	fmt.Println(status)

	// it's possible to convert to float
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(grade)
}
