package main

import "fmt"

func main() {
	fmt.Println("1/3: ", 1.0/3.0) // 0.3333333333333333
	// output is ugly. round is needed
	fmt.Printf("1/3: %.2f\n", 1.0/3.0)
	// Printf doesn't add \n automatically
	// %s string
	// %d int
	// %f float
	fmt.Printf("The %s const %d cents each.\n", "gumballs", 23)
	// print string with max length 12
	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("--------------")
	// print int with max length 2
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	// %5.3f  5 - length of the whole num. 3 - max length of the second part. 12.345
	// %.2f - no max length of int!
}
