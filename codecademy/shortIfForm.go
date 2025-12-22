package main

import "fmt"

func main() {

	// short-cut if. scope of success is if block
	if success := true; success {
		fmt.Println("We're rich!")
	} else {
		fmt.Println("Where did we go wrong?")
	}

	amountStolen := 50000
	// short-cut switch. scope of numOfThieves is switch block
	switch numOfThieves := 5; numOfThieves {
	case 1:
		fmt.Println("I'll take all $", amountStolen)
	case 2:
		fmt.Println("Everyone gets $", amountStolen/2)
	case 3:
		fmt.Println("Everyone gets $", amountStolen/3)
	case 4:
		fmt.Println("Everyone gets $", amountStolen/4)
	case 5:
		fmt.Println("Everyone gets $", amountStolen/5)
	default:
		fmt.Println("There's not enough to go around...")
	}
}
