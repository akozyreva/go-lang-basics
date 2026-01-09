package main

import "fmt"

type subscriber struct {
	name   string
	rate   float32
	active bool
}

func showInfo(person *subscriber) {
	fmt.Printf("Name: %s\n", person.name)
	fmt.Printf("Rate: %f\n", person.rate)
	fmt.Printf("Active: %t\n", person.active)
}

func defaultSubscriber(name string) subscriber {
	return subscriber{
		name:   name,
		rate:   0.0,
		active: true,
	}
}

func main() {
	var student struct {
		name string
		rate float64
	}
	student.name = "Alice"
	student.rate = 0.1
	fmt.Println(student.rate)

	subscriber1 := subscriber{
		name:   "Bob",
		rate:   0.1,
		active: true,
	}
	subscriber2 := subscriber{
		name:   "John",
		rate:   0.2,
		active: false,
	}
	fmt.Println(subscriber1, subscriber2)
	showInfo(&subscriber1)
	subscriber3 := defaultSubscriber("Jack")
	showInfo(&subscriber3)
}
