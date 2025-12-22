package main

import (
	"fmt"
	"time"
)

func main() {
	var notes [7]string
	fmt.Println(notes)
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	// the rest values are set to ''
	for index, note := range notes {
		fmt.Println(index, note)
	}
	// if you don't need to index
	for _, note := range notes {
		fmt.Println(note)
	}

	var primes [5]int
	primes[0] = 2
	primes[1] = 3
	// the rest values are set to 0
	fmt.Println(primes)

	// arr of Times!
	var dates [3]time.Time
	dates[0] = time.Now()
	dates[1] = time.Now().Add(time.Hour)
	// def value of Time is 0001-01-01 00:00:00 +0000 UTC
	fmt.Println(dates)

	// init arr on the fly
	var days = [7]int{1, 2, 3, 4, 5, 6, 7}
	// or:
	weekDays := [7]string{"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday", // but if leave "Sunday"}, comma isn't required!!
	}
	fmt.Println(days)
	fmt.Println(weekDays)
	// Printf for printing arr
	fmt.Printf("%v\n", weekDays)
	// len of arr
	fmt.Println(len(weekDays))

	kgOfMeats := [3]float64{71.0, 56.2, 89.9}
	sum := 0.0
	countOfWeeks := len(kgOfMeats)
	for _, kgOfMeat := range kgOfMeats {
		sum += kgOfMeat
	}
	fmt.Printf("%.2f", sum/float64(countOfWeeks))
}
