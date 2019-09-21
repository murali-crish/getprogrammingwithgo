package main

import "fmt"

func main() {
	var command = "go east"

	if command == "go east" {
		fmt.Println("You head further up the mountain.")
	} else if command == "go inside" {
		fmt.Println("You enter the cave where you live out the rest of life.")
	} else {
		fmt.Println("Didn't quite get that.")
	}

	var room = "building"
	if room == "cave" {
		fmt.Println("It's a cave.")
	} else if room == "entrance" {
		fmt.Println("It's an entrance.")
	} else if room == "mountain" {
		fmt.Println("It's mountain.")
	} else {
		fmt.Println("It's not known.")
	}
}
