package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	// math/rand - unique path to package. by convention rand is the name to call methods in package
	"math/rand"
)

func main() {
	var countsOfAttempts int = 10
	var isGuessedNum bool = false
	seconds := time.Now().Unix() // current time in int (Unix())
	fmt.Println(seconds)
	// rand.Seed is needed to generate real random number
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println(target)
	fmt.Println("I've chosen a random num between 1 and 100")
	fmt.Println("Can you guess it?")
	i := countsOfAttempts
	// it's short form. full is for i := countsOfAttempts; i > 0; i--
	for i > 0 {
		fmt.Println("You have :", i, "attempts")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		inputNum, err := strconv.Atoi(input) // convert to int
		if err != nil {
			log.Fatal(err)
		}
		i--
		if inputNum == target {
			fmt.Println("Your guess was correct. End")
			isGuessedNum = true
			break
		} else {
			fmt.Println("Ooops.. try again")
			if inputNum > target {
				fmt.Println("Your guess was HIGH")
			} else {
				fmt.Println("Your guess was LOW")
			}
		}
	}
	if !isGuessedNum {
		fmt.Println("You lost. Correct number is: ", target)
	}

}
