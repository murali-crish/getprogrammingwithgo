package main

import (
	"fmt"
	"math/rand"
)

func main() {
	piggyBank := 0.0
	currency := ""
	for i := 0; piggyBank < 20.00; i++ {
		switch rand.Intn(3) {
		case 0:
			piggyBank += 0.05 // nickels
			currency = "nickel"
		case 1:
			piggyBank += 0.10 // dimes
			currency = "dime"
		case 2:
			piggyBank += 0.25 //quarters
			currency = "quarter"
		}
		fmt.Printf("Adding a %s - %.2f\n", currency, piggyBank)
	}
	fmt.Println(piggyBank)
}
