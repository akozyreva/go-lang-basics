package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type NoiseMaker interface {
	MakeSound()
}

// PlayWithToy one function for all types
func PlayWithToy(toy NoiseMaker) {
	fmt.Println("Playing with toy:")
	toy.MakeSound()
}

func main() {
	var toy NoiseMaker // 1 variable!
	toy = Whistle("Toyco Canary")
	toy.MakeSound()
	toy = Horn("Toyco Blaster") // because you're not allowed to reuse vars, but with interface you can
	toy.MakeSound()

	// arr of different types!
	toys := []NoiseMaker{
		Whistle("Canary"),
		Horn("Blaster"),
		Whistle("Robin"),
		Horn("Trumpet"),
	}

	// one cycle can handle it
	for _, toy := range toys {
		toy.MakeSound()
	}

	PlayWithToy(Whistle("Toyco Canary"))
}
