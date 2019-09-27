package main

import (
	"fmt"
	"math/rand"
)

func main() {
	spaceline := ""
	fmt.Printf("%v\t\t %v\t %v\t %v\t %v\n", "Spaceline", "Days", "Trip", "type", "Price")
	fmt.Println("==============================================")
	for i := 0; i < 10; i++ {
		speed := rand.Intn(15) + 16
		day := rand.Intn(100) + 1
		// price := rand.Intn(100) + 1
		switch rand.Intn(3) {
		case 0:
			spaceline = "Space Adventures"
		case 1:
			spaceline = "SpaceX"
		case 2:
			spaceline = "Virgin Glactic"
		}
		fmt.Printf("%v\t\t %v\t %v\t %v\t %v\n", spaceline, day, "Trip", "type", speed)
	}
}
