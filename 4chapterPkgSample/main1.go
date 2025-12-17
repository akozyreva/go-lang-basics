package main

import (
	"4chapterPkgSample/dates"
	"4chapterPkgSample/greeting"
	// sample of using nested pkg
	"4chapterPkgSample/greeting/deutsch"
	"fmt"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	deutsch.Hallo()
	// import constant
	fmt.Println(dates.DaysInWeek)
}
