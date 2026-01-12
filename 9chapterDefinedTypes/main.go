package main

import (
	"fmt"
	"time"
)

// Liters usually named types are structs, but
// basic types also can be used
type Liters float64 // it's ambigiously, because you can use there simply small variable. but it can be used for better code reading
type Gallons float64
type Seconds int

func setTimeout(duration Seconds) {
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Println("Timeout completed!")
}

func ToGallons(liters Liters) Gallons {
	return Gallons(liters * 0.264)
}

func ToLiters(gallons Gallons) Liters {
	return Liters(gallons * 3.785)
}
func main() {
	// Вместо:
	//setTimeout(5000) // Это секунды или миллисекунды? 🤔

	// Явно:
	setTimeout(Seconds(5)) // Понятно!

	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(10.0)
	busFuel = Liters(240.0)
	fmt.Println(carFuel, busFuel)
	// to convert one type to another, you can do it only like this...
	carFuel = Gallons(Liters(40) * 0.264)
	fmt.Println(carFuel)
	// you can use operators here, which supports your basic type
	carFuel += Gallons(1)
	carFuel += 2
	fmt.Println(carFuel)
	carFuel += ToGallons(Liters(10))
	fmt.Println(carFuel)
}
