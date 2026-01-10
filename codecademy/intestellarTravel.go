package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
	fmt.Println("Fuel is left: ", fuel)
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
	var fuel int
	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	}
	return fuel
}

// Create the function greetPlanet() here

// Create the function cantFly() here

// Create the function flyToPlanet() here

func main() {
	// Test your functions!
	fuelGauge(120)
	fuel := calculateFuel("Venus")
	fmt.Println(fuel)
	// Create `planetChoice` and `fuel`

	// And then liftoff!

}
