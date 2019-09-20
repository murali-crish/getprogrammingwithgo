package main

import "fmt"

func main() {
	const hoursPerDay = 24
	var days, distance = 28, 56000000
	fmt.Printf("Ship should travel at %v kms/h to reach Malacandra in %v days\n", distance/days/hoursPerDay, days)
}
