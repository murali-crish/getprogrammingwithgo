// How long does it take to get to Mars?
package main

import "fmt"

func main() {
	const hoursPerDay = 24
	var (
		speed    = 100800   // km/s
		distance = 56000000 // km
	)

	fmt.Println(distance/speed/hoursPerDay, "seconds")
}
