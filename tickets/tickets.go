package main

import (
	"fmt"
	"math/rand"
)

const secondsPerDay = 86400

func main() {
	distance := 62100000
	spaceline := ""

	fmt.Printf("%-16v %4v %-8v %v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("========================================================")
	for i := 0; i < 10; i++ {
		// price := rand.Intn(100) + 1
		switch rand.Intn(3) {
		case 0:
			spaceline = "Space Adventures"
		case 1:
			spaceline = "SpaceX"
		case 2:
			spaceline = "Virgin Glactic"
		}

		speed := rand.Intn(15) + 16
		duration := distance / speed / secondsPerDay
		price := 20.0 + speed
		trip := "One-way"

		if rand.Intn(2) == 1 {
			trip = "Round-trip"
			price = price * 2
		}

		fmt.Printf("%-16v %4v %-8v $%v\n", spaceline, duration, trip, price)
	}
}
