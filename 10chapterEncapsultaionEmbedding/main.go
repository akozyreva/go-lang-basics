package main

import (
	"10chapterEncapsultaionEmbedding/calendar"
	"fmt"
	"log"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2000)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(1)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(12)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
	// the same this code doesn't work: fmt.Println(date.day)
	// use getter instead
	year := date.Year()
	fmt.Println(year)
	date1 := calendar.Date{}
	// date1.day = 1 // it doesn't work anymore, because params inside of struct are not exported!
	fmt.Println(date1)

	event := calendar.Event{}
	event.SetTitle("Concert")
	err = event.SetDay(1) // event 'inherited' all Date methods...
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2027)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event)

}
