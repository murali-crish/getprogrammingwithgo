package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var myPick = 67
	var randomNumber = rand.Intn(100) + 1
	for randomNumber != myPick {
		if randomNumber < myPick {
			fmt.Println("Too small, not ", randomNumber)
			randomNumber = rand.Intn(100-randomNumber) + 1
		} else {
			fmt.Println("Too Big, not ", randomNumber)
			randomNumber = rand.Intn(randomNumber) + 1
		}
	}
	fmt.Println("Picked my number")

}
